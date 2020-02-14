package frontend

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type IndexController struct {
	Ctx iris.Context
}

func (r *IndexController) Get() mvc.View {
	return mvc.View{
		Name:   "frontend/index/index.html",
		Layout: "shared/layoutFront.html",
		Data: iris.Map{
			"Title":   "首页",
			"Message": "Message 成功了",
		},
	}
}
