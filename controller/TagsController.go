package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	// commons "iris/commons"
	"iris/model"
	"log"
)

// var TagsModel = model.Tags{}

type TagsController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *TagsController) Get() mvc.View {
	Tags := model.Tags{}
	list := Tags.List()
	log.Println(list)
	return mvc.View{
		Name: "tags/list.html",
		Data: iris.Map{
			"Title": "标签列表",
			"list":  list,
		},
	}
}
