package frontend

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris/commons"

	// "github.com/kataras/iris/sessions"
	"iris/model"
	"strconv"
	"strings"
)

type IndexController struct {
	Ctx iris.Context
	News model.News
	// Session *sessions.Sessions
}

func (r *IndexController) Get() mvc.View {
	page, err := strconv.Atoi(r.Ctx.URLParam("page"))

	if err != nil || page < 1 {
		page = 1
	}

	list, total, totalPages := r.News.List(page)

	Category := model.Category{}
	Tag := model.Tags{}

	for k, v := range list {
		CategoryName := ""
		if val, err := Category.CategoryMoreInfo(v.Category_id); err == nil {
			for _, vv := range val {
				CategoryName += vv.Name + ","
			}
		}
		list[k].CategoryName = strings.TrimRight(CategoryName, ",")
	}
	for k, v := range list {
		TagsName := ""
		if val, err := Tag.TagsMoreInfo(v.Tags_id); err == nil {
			for _, vv := range val {
				TagsName += vv.Name + ","
			}
		}
		list[k].TagsName = strings.TrimRight(TagsName, ",")
	}
	// log.Println(list, total, totalPages)
	return mvc.View{
		Name:   "frontend/index/index.html",
		Layout: "shared/layoutFront.html",
		Data: iris.Map{
			"Title":   "首页",
			"Message": "Message 成功了 嘻嘻",
			"list": list,
			"PageHtml": commons.GetPageHtml(totalPages, page, total, r.Ctx.Path()),
		},
	}
}
