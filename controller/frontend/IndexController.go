package frontend

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris/commons"
	"iris/libs/logging"

	// "github.com/kataras/iris/sessions"
	"log"
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
	NewsNewest := r.News.NewsNewest()
	Category := model.Category{}
	Tag := model.Tags{}
	CategoryList := Category.ListFrontend()
	TagList := Tag.ListAll()
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
	// log.Println(CategoryList)
	return mvc.View{
		Name:   "frontend/index/index.html",
		Layout: "shared/layoutFront.html",
		Data: iris.Map{
			"Title":   "首页",
			"Message": "Message 成功了 嘻嘻",
			"list": list,
			"PageHtml": commons.GetPageHtml(totalPages, page, total, r.Ctx.Path()),
			"CategoryList": CategoryList,
			"TagList": TagList,
			"NewsNewest": NewsNewest,
		},
	}
}

func (r *IndexController) GetSearch() mvc.View {
	search := r.Ctx.URLParam("q")
	log.Println(search)
	searchList := r.News.NewsSearch(search)
	Category := model.Category{}
	Tag := model.Tags{}
	for k, v := range searchList {
		CategoryName := ""
		if val, err := Category.CategoryMoreInfo(v.Category_id); err == nil {
			for _, vv := range val {
				CategoryName += vv.Name + ","
			}
		}
		searchList[k].CategoryName = strings.TrimRight(CategoryName, ",")
	}
	for k, v := range searchList {
		TagsName := ""
		if val, err := Tag.TagsMoreInfo(v.Tags_id); err == nil {
			for _, vv := range val {
				TagsName += vv.Name + ","
			}
		}
		searchList[k].TagsName = strings.TrimRight(TagsName, ",")
	}
	CategoryList := Category.ListFrontend()
	TagList := Tag.ListAll()
	NewsNewest := r.News.NewsNewest()
	return mvc.View{
		Name:   "frontend/index/index.html",
		Layout: "shared/layoutFront.html",
		Data: iris.Map{
			"Title":   "查询结果",
			"Message": "Message 成功了 嘻嘻",
			"list": searchList,
			//"PageHtml": commons.GetPageHtml(totalPages, page, total, r.Ctx.Path()),
			"CategoryList": CategoryList,
			"TagList": TagList,
			"NewsNewest": NewsNewest,
		},
	}
}

func (r *IndexController) GetAbout() mvc.View{
	return mvc.View{
		Name: "frontend/about.html",
		Layout: "shared/layoutFront.html",
	}
}

func (r *IndexController) GetArticle() mvc.View {
	page, err := strconv.Atoi(r.Ctx.URLParam("page"))

	if err != nil || page < 1 {
		page = 1
	}

	list, total, totalPages := r.News.List(page)
	log.Println(list)
	// NewsNewest := r.News.NewsNewest()
	Category := model.Category{}
	Tag := model.Tags{}
	// CategoryList := Category.ListFrontend()
	// TagList := Tag.ListAll()
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
		Name: "frontend/full-width.html",
		Layout: "shared/layoutFront.html",
		Data: iris.Map{
			"Title":   "博客",
			"list": list,
			"PageHtml": commons.GetPageHtml(totalPages, page, total, r.Ctx.Path()),
			//"CategoryList": CategoryList,
			//"TagList": TagList,
			//"NewsNewest": NewsNewest,
		},
	}
}

func (r *IndexController) GetArticleBy(id uint) mvc.View {
	info, err := r.News.NewsInfo(id)
	if err != nil {
		logging.Info(err)
	}

	log.Println(info.Content)
	CategoryIds := []string{}
	for _, v := range strings.Split(info.Category_id, ",") {
		_v, _ := strconv.Atoi(v)
		var categoryInfo model.Category
		categoryDetail, err := categoryInfo.CategoryInfo(uint(_v))
		if err != nil {
			logging.Info(err)
		}
		CategoryIds = append(CategoryIds, categoryDetail.Name)
	}

	tagIds := []string{}
	// if NewsInfo.Tags_id
	for _, v := range strings.Split(info.Tags_id, ",") {
		_v, _ := strconv.Atoi(v)
		var tagsInfo model.Tags
		tagsDetail, err := tagsInfo.TagInfo(uint(_v))
		if err != nil {
			logging.Info(err)
		}
		tagIds = append(tagIds, tagsDetail.Name)
	}


	return mvc.View{
		Name: "frontend/single.html",
		Layout: "shared/layoutFront.html",
		Data: iris.Map{
			"Title": "文章详情",
			"Info": info,
			"CategoryIds": CategoryIds,
			"tagIds": tagIds,
		},
	}
}