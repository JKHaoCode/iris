package route

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris/commons"
	"iris/controller"
	"iris/middleware"
)

func Routes(app *iris.Application) {
	//登录路由
	mvc.New(app.Party("/")). // 根据请求类型和请求URL自动匹配处理方法 contoller 方法
					Register(commons.SessManager.Start).
					Handle(new(controllers.IndexController))

	//登录路由
	mvc.New(app.Party("/login")).
		Register(commons.SessManager.Start).
		Handle(new(controllers.LoginController))

	//系统路由
	mvc.New(app.Party("/backend/system", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.SystemController))
	//管理员管理
	mvc.New(app.Party("/backend/administrators", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.AdministratorsController))
	//分类管理
	mvc.New(app.Party("/backend/categorys", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.CategorysController))
	//内容管理
	mvc.New(app.Party("/backend/news", middleware.SessionLoginAuth)).
		Register(commons.SessManager.Start).
		Handle(new(controllers.NewsController))
}
