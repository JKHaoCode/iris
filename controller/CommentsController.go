package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris/commons"
	"iris/model"
	"log"
	"strconv"
)

type CommentsController struct {
	Ctx          iris.Context
	CommentModel model.Comment
}

func (c *CommentsController) Get() mvc.View {
	page, err := strconv.Atoi(c.Ctx.URLParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	var artile model.News
	comment := model.Comment{}
	list, total, totalPages := comment.List(page)
	// log.Println(list)

	for k, v := range list {
		ArticleName := ""
		if val, err := artile.NewsInfo(v.ArticleId); err == nil {
			ArticleName += val.Title
		}
		list[k].ArticleName = ArticleName
	}
	return mvc.View{
		Name: "comment/list.html",
		Data: iris.Map{
			"Title":    "评论列表",
			"list":     list,
			"PageHtml": commons.GetPageHtml(totalPages, page, total, c.Ctx.Path()),
		},
	}
}
