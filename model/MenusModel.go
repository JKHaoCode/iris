package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"iris/libs"
	"log"
	"strings"
)

var ListMenusTree []Menus
var ListMenusToTree []TreeList

type Menus struct {
	gorm.Model
	Name     string `gorm:"not null;VARCHAR(100);"validate:"required"`
	URL 	 string `gorm:"not null;VARCHAR(100);"validate:"required"`
	Icon	 string `gorm:"VARCHAR(50)"`
	ParentId int    `gorm:"default:'0';not null;"`
	Sort     int    `gorm:"default:'0';not null;"validate:"number,min=0"`
	Level 	 int 	`gorm:"-"`
}

type TreeList struct {
	gorm.Model
	Name     string `gorm:"not null;VARCHAR(100);"validate:"required"`
	URL 	 string `gorm:"not null;VARCHAR(100);"validate:"required"`
	Icon	 string `gorm:"VARCHAR(50)"`
	ParentId int    `gorm:"default:'0';not null;"`
	Sort     int    `gorm:"default:'0';not null;"validate:"number,min=0"`
	Level 	 int 	`gorm:"-"`
	Children []TreeList	`gorm:"-"`
}

func (this *Menus) List() []Menus {
	var data = []Menus{}
	db := libs.DB

	err := db.Order("sort").Find(&data).Error
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func (this *Menus) MenusInfo(id uint) (Menus, error) {
	var menus Menus
	db := libs.DB

	if db.Where("id = ?", id).First(&menus).RecordNotFound() {
		return Menus{}, errors.New("分类未找到")
	}
	return menus, nil
}

func (this *Menus) MenusMoreInfo(ids string) ([]Menus, error) {
	var data = []Menus{}
	db := libs.DB

	if db.Where("id in (?)", strings.Split(ids, ",")).Find(&data).RecordNotFound() {
		return []Menus{}, errors.New("分类未找到[]")
	}
	return data, nil
}

func (this *Menus) MenusAdd(postValues map[string][]string) error {
	var menus Menus
	db := libs.DB

	if err := libs.FormDecode(&menus, postValues); err != nil {
		return err
	}
	if err := libs.Validate(menus); err != nil {
		return err
	}
	if !db.Where("name = ? ", menus.Name).First(&Menus{}).RecordNotFound() {
		return errors.New("该名称已经存在")
	}
	if err := db.Create(&menus).Error; err != nil {
		return err
	}
	return nil
}

func (this *Menus) MenusUpdate(postValues map[string][]string) error {
	// log.Println(postValues)
	var menus Menus
	db := libs.DB

	if err := libs.FormDecode(&menus, postValues); err != nil {
		return err
	}
	if err := libs.Validate(menus); err != nil {
		return err
	}
	if !db.Where("name = ? and id != ?", menus.Name, menus.ID).Find(&Menus{}).RecordNotFound() {
		return errors.New("该名称已经存在")
	}
	if db.Where("id = ? ", menus.ID).Find(&Menus{}).RecordNotFound() {
		return errors.New("未查询到分类id")
	}
	if err := db.Save(&menus).Error; err != nil {
		return err
	}
	return nil
}

func (this *Menus) MenusDel(id uint) error {
	var menus Menus
	db := libs.DB

	if !db.Where("parent_id = ?", id).Find(&menus).RecordNotFound() {
		return errors.New("该分类下存在子级分类，请先删除子级分类")
	}
	if err := db.Where("id = ?", id).Delete(&menus).Error; err != nil {
		return err
	}
	return nil
}

func (this *Menus) GetTree(list []Menus, pid int, level int) []Menus {
	for _, v := range list {
		if v.ParentId == pid {
			v.Level = level
			v.Name = strings.Repeat("————", v.Level) + v.Name
			ListMenusTree = append(ListMenusTree, v)
			/*if len(list) == 0 {
				list = []Category{}
			} else {
				list = append(list[:index], list[index+1:]...)
			}*/
			this.GetTree(list, int(v.Model.ID), v.Level+1)
		}
	}
	return ListMenusTree
}

func (m *Menus)GetMenu(menu []Menus, pid int) []TreeList {
	treeList := []TreeList{}
	for _, v := range menu{
		if v.ParentId == pid {
			child := v.GetMenu(menu, int(v.ID))
			node := TreeList{
				// ID:       v.ID,
				Name:     v.Name,
				Sort:     v.Sort,
				URL:      v.URL,
				ParentId: v.ParentId,
				Icon: v.Icon,
			}
			node.Children = child
			treeList = append(treeList, node)
		}
	}
	// log.Println(treeList)
	return treeList
}