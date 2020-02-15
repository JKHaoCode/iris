package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris/commons"
	"iris/model"
	"log"
	"strconv"
)

var TagsModel = model.Tags{}

type TagsController struct {
	Ctx     iris.Context
	// Session *sessions.Session
	Tags    model.Tags
}

func (c *TagsController) Get() mvc.View {
	page, err := strconv.Atoi(c.Ctx.URLParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	// Tags := model.Tags{}
	list, total, totalPages := c.Tags.List(page)
	log.Println(list)
	return mvc.View{
		Name: "tags/list.html",
		Data: iris.Map{
			"Title":    "标签列表",
			"list":     list,
			"PageHtml": commons.GetPageHtml(totalPages, page, total, c.Ctx.Path()),
		},
	}
}

func (c *TagsController) GetAddTag() mvc.View {
	return mvc.View{
		Name: "tags/addTag.html",
		Data: iris.Map{
			"Title": "新增标签",
		},
	}
}

func (c *TagsController) GetUpdateTagBy(id uint) mvc.View {
	tag, err := TagsModel.TagInfo(id)
	if err != nil {
		return commons.MvcError(err.Error(), c.Ctx)
	}

	return mvc.View{
		Name: "tags/updateTag.html",
		Data: iris.Map{
			"Title":         "标签修改",
			"UpdateTagInfo": tag,
		},
	}
}

func (c *TagsController) PostAddTag() {
	if err := TagsModel.TagsAdd(c.Ctx.FormValues()); err == nil {
		c.Ctx.Redirect("/backend/tags")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *TagsController) PostUpdateTag() {
	if err := TagsModel.TagsUpdate(c.Ctx.FormValues()); err == nil {
		c.Ctx.Redirect("/backend/tags")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *TagsController) GetDelTagBy(id uint) {
	if err := TagsModel.TagDel(id); err == nil {
		c.Ctx.Redirect("/backend/tags")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}
