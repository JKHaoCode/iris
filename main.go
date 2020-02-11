package main

import (
	"./commons"
	"./libs"
	"./model"
	"./route"
	iris "github.com/kataras/iris"
	context "github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	config "github.com/spf13/viper"
	"log"
	"strconv"
	"time"
)

func init() {
	config.AddConfigPath("./configs")
	config.SetConfigName("mysql")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	dbConfig := libs.DbConfig{
		config.GetString("default.host"),
		config.GetString("default.port"),
		config.GetString("default.database"),
		config.GetString("default.user"),
		config.GetString("default.password"),
		config.GetString("default.charset"),
		config.GetInt("default.MaxIdleConns"),
		config.GetInt("default.MaxOpenConns"),
	}
	libs.DB = dbConfig.InitDB()
	if config.GetBool("default.sql_log") {
		libs.DB.LogMode(true)
	}
}

func main() {
	app := iris.New()
	config.SetConfigName("app")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件错误, %s", err)
	}
	tmpl := iris.HTML("./views", ".html").Layout(config.GetString("site.DefaultLayout"))
	if config.GetBool("site.APPDebug") == true {
		app.Logger().SetLevel("debug") //设置debug
		tmpl.Reload(true)
	}

	tmpl.AddFunc("TimeToDate", libs.TimeToDate)
	tmpl.AddFunc("strToHtml", libs.StrToHtml)

	app.RegisterView(tmpl)
	app.Favicon("./favicon.ico")
	app.Use(iris.Gzip)

	//（可选）添加两个内置处理程序
	//可以从任何与http相关的panics中恢复
	//并将请求记录到终端。
	app.Use(recover.New())
	app.Use(logger.New())

	app.StaticWeb("/public", "./public")   //设置静态文件目录
	app.StaticWeb("/uploads", "./uploads") //设置静态文件目录
	//设置公共页面输出
	app.Use(func(ctx iris.Context) {
		if auth := commons.SessManager.Start(ctx).Get("admin_user"); auth != nil {
			admin_user, _ := auth.(map[string]interface{})
			var admin_model model.Admin
			admin_id, _ := admin_user["id"].(uint)
			adminInfo, _ := admin_model.AdminInfo(admin_id)
			if adminInfo.Headico == "" {
				adminInfo.Headico = "/public/adminlit/dist/img/user2-160x160.jpg"
			}
			ctx.ViewData("adminInfo", adminInfo)
		}
		ctx.ViewData("Title", config.GetString("site.DefaultTitle"))
		now := time.Now().Format(ctx.Application().ConfigurationReadOnly().GetTimeFormat())
		ctx.ViewData("CurrentTime", now)
		ctx.Next()
	})

	//设置错误模版
	app.OnAnyErrorCode(func(ctx iris.Context) {
		_, err := ctx.HTML("<center>很抱歉！当前页面错误,错误代码:" + strconv.Itoa(ctx.GetStatusCode()) + "</center>")
		if err != nil {
			log.Fatalf("内部错误,错误代码 %s", err)
		}
	})

	route.Routes(app)

	//应用配置文件
	app.Configure(iris.WithConfiguration(iris.YAML("./configs/iris.yml")))

	//Run
	www := app.Party("www.")
	{
		currentRoutes := app.GetRoutes()
		for _, r := range currentRoutes {
			www.Handle(r.Method, r.Tmpl().Src, r.Handlers...)
		}
	}
	err := app.Run(iris.Addr(config.GetString("server.domain") + ":" + config.GetString("server.port")))
	if err != nil {
		log.Fatalf("服务启动失败,错误代码 %s", err)
	}
}
