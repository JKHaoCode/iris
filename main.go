package main

import (
	"github.com/kataras/iris"
	// "iris/libs/logging"
	"iris/libs/redis"

	// context "github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	config "github.com/spf13/viper"
	"iris/commons"
	"iris/libs"
	"iris/model"
	"iris/route"
	"log"
	// "strconv"
	"time"
	// "os"
	// "os/signal"
	// "fmt"
)

func init() {
	config.AddConfigPath("./configs")
	config.SetConfigName("mysql") // config file
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
	redis.Setup()
	//redis := redis2.Singleton()
	//commons.SessManager.UseDatabase(redis)
}

func main() {
	// fmt.Printf("%T\n", config.GetInt64("site.SessionExpires"))
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("/frontend")
	}) // 访问 / 时 自动跳转 /frontend
	app.Get("/backend", func(ctx iris.Context) {
		ctx.Redirect("/backend/system")
	})
	app.Get("/backend/system", func(ctx iris.Context) {
		ctx.Redirect("/backend/system/main")
	})
	config.SetConfigName("app") //读取的文件
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件错误, %s", err)
	}
	tmpl := iris.HTML("./views", ".html").Layout(config.GetString("site.DefaultLayout"))
	if config.GetBool("site.APPDebug") == true {
		app.Logger().SetLevel("debug") //设置debug
		tmpl.Reload(true)
	}

	tmpl.AddFunc("TimeToDate", libs.TimeToDate) // 为html 页面增加func
	tmpl.AddFunc("strToHtml", libs.StrToHtml)   // 为html 页面增加func 用法{{.created_at | func}}
	tmpl.AddFunc("AddKey", libs.AddKey)
	tmpl.AddFunc("timeNow", libs.TimeYear)

	app.RegisterView(tmpl)
	app.Favicon("./favicon.ico")
	app.Use(iris.Gzip)

	//（可选）添加两个内置处理程序
	//可以从任何与http相关的panics中恢复
	//并将请求记录到终端。
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(libs.Cors)

	app.StaticWeb("/public", "./public")   //设置静态文件目录
	app.StaticWeb("/uploads", "./uploads") //设置静态文件目录
	// 设置公共页面输出 重点
	app.Use(func(ctx iris.Context) {
		if auth := commons.SessManager.Start(ctx).Get("admin_user"); auth != nil {
			admin_user, _ := auth.(map[string]interface{})
			var admin_model model.Admin
			var menu model.Menus
			list := menu.List()
			listToTree := menu.GetMenu(list, 999)
			admin_id, _ := admin_user["id"].(uint)
			adminInfo, _ := admin_model.AdminInfo(admin_id)
			// log.Println(adminInfo)
			if adminInfo.Headico == "" {
				adminInfo.Headico = "/public/adminlit/dist/img/user2-160x160.jpg"
			}
			// log.Println(adminInfo)
			ctx.ViewData("adminInfo", adminInfo)
			ctx.ViewData("listToTree", listToTree)
		}
		ctx.ViewData("Title", config.GetString("site.DefaultTitle"))
		now := time.Now().Format(ctx.Application().ConfigurationReadOnly().GetTimeFormat())
		ctx.ViewData("CurrentTime", now)
		ctx.Next()
	})

	//设置错误模版
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewLayout("shared/layoutNone.html")
		ctx.ViewData("Message", 404)
		ctx.View("errors/404.html")
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
	// go func() {
	// 	err := app.Run(iris.Addr(config.GetString("server.domain") + ":" + config.GetString("server.port")))
	// 	if err != nil {
	// 		log.Fatalf("服务启动失败,错误代码 %s", err)
	// 	}
	// }()

	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// log.Println("closing database connection")
	// eng
}
