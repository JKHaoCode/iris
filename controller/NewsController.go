package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris/commons"
	"iris/model"
	"strconv"
	"strings"
)

type NewsController struct {
	Ctx     iris.Context
	// Session *sessions.Session
	News    model.News
}

func (c *NewsController) Get() mvc.View {
	page, err := strconv.Atoi(c.Ctx.URLParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	list, total, totalPages := c.News.List(page)
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
	return mvc.View{
		Name: "news/list.html",
		Data: iris.Map{
			"Title":    "内容列表",
			"list":     list,
			"PageHtml": commons.GetPageHtml(totalPages, page, total, c.Ctx.Path()),
		},
	}
}

func (c *NewsController) GetAddNews() mvc.View {
	Category := model.Category{}
	Tag := model.Tags{}
	list := Category.List()
	tagList := Tag.ListAll()
	model.ListTree = []model.Category{}
	list = Category.GetTree(list, 0, 0)
	return mvc.View{
		Name: "news/addNews.html",
		Data: iris.Map{
			"Title":   "新增内容",
			"list":    list,
			"tagList": tagList,
		},
	}
}

func (c *NewsController) PostAddNews() {
	if err := c.News.NewsAdd(c.Ctx.FormValues()); err == nil {
		c.Ctx.Redirect("/backend/news")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *NewsController) GetUpdateNewsBy(id uint) mvc.View {
	NewsInfo, err := c.News.NewsInfo(id)
	if err != nil {
		return commons.MvcError(err.Error(), c.Ctx)
	}
	Category := model.Category{}
	list := Category.List()
	model.ListTree = []model.Category{}
	list = Category.GetTree(list, 0, 0)

	CategoryIds := []int{}
	for _, v := range strings.Split(NewsInfo.Category_id, ",") {
		_v, _ := strconv.Atoi(v)
		CategoryIds = append(CategoryIds, _v)
	}

	Tag := model.Tags{}
	tagList := Tag.ListAll()
	tagIds := []int{}
	// if NewsInfo.Tags_id
	for _, v := range strings.Split(NewsInfo.Tags_id, ",") {
		_v, _ := strconv.Atoi(v)
		tagIds = append(tagIds, _v)
	}

	return mvc.View{
		Name: "news/updateNews.html",
		Data: iris.Map{
			"Title":          "内容修改",
			"UpdateNewsInfo": NewsInfo,
			"CategoryIds":    CategoryIds,
			"list":           list,
			"tagList":        tagList,
			"tagIds":         tagIds,
		},
	}
}

func (c *NewsController) PostUpdateNews() {
	if err := c.News.NewsUpdate(c.Ctx.FormValues()); err == nil {
		c.Ctx.Redirect("/backend/news")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *NewsController) GetDelNewsBy(id uint) {
	if err := c.News.NewsDel(id); err == nil {
		c.Ctx.Redirect("/backend/news")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}
