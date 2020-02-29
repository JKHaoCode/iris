package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris/commons"
	"iris/model"

	"log"
)

var MenusModel = model.Menus{}

type MenusController struct {
	Ctx     iris.Context
	// Session *sessions.Session
}

func (c *MenusController) Get() mvc.View {
	Menus := model.Menus{}
	list := Menus.List()
	model.ListMenusTree = []model.Menus{}
	list = Menus.GetTree(list, 0, 0)
	return mvc.View{
		Name: "menus/list.html",
		Data: iris.Map{
			"Title": "菜单列表",
			"list":  list,
		},
	}
}

func (c *MenusController) GetAddMenu() mvc.View {
	Menu := model.Menus{}
	list := Menu.List()
	model.ListMenusTree = []model.Menus{}
	list = Menu.GetTree(list, 999, 0)
	return mvc.View{
		Name: "menus/addMenus.html",
		Data: iris.Map{
			"Title": "新增菜单",
			"list":  list,
		},
	}
}

func (c *MenusController) PostAddMenu() {
	if err := MenusModel.MenusAdd(c.Ctx.FormValues()); err == nil {
		c.Ctx.Redirect("/backend/menus")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *MenusController) GetUpdateMenuBy(id uint) mvc.View {
	menuInfo, err := MenusModel.MenusInfo(id)
	if err != nil {
		return commons.MvcError(err.Error(), c.Ctx)
	}
	Menu := model.Menus{}
	list := Menu.List()
	log.Println(list)
	model.ListMenusTree = []model.Menus{}
	list = Menu.GetTree(list, 0, 0)

	return mvc.View{
		Name: "menus/updateMenus.html",
		Data: iris.Map{
			"Title":          "菜单修改",
			"UpdateMenuInfo": menuInfo,
			"list":           list,
		},
	}
}

func (c *MenusController) PostUpdateMenu() {
	// log.Println(c.Ctx.FormValues())
	if err := MenusModel.MenusUpdate(c.Ctx.FormValues()); err == nil {
		c.Ctx.Redirect("/backend/menus")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *MenusController) GetDelCategoryBy(id uint) {
	if err := CategoryModel.CategoryDel(id); err == nil {
		c.Ctx.Redirect("/backend/menus")
	} else {
		commons.DefaultErrorShow(err.Error(), c.Ctx)
	}
}
