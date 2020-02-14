package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	config "github.com/spf13/viper"
	"iris/libs"
	"log"
	"math"
	"strings"
)

// var ListTree []Category

type Tags struct {
	gorm.Model
	Name string `gorm:"not null;VARCHAR(100);"validate:"required"`
	Sort int    `gorm:"default:'0';not null;"validate:"number,min=0"`
	// Level int    `gorm:"-"`
}

func (this *Tags) List(page int) ([]Tags, int, int) {
	var data = []Tags{}
	var totalCount int
	db := libs.DB

	err := db.Order("sort desc").Find(&data).Error
	if err != nil {
		log.Fatalln(err)
	}

	limit := config.GetInt("pagination.PageSize")
	offset := (page - 1) * limit
	db.Find(&data).Count(&totalCount)
	db.Offset(offset).Limit(limit).Order("id desc").Find(&data)
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return data, totalCount, totalPages
}

func (this *Tags) TagInfo(id uint) (Tags, error) {
	var tag Tags
	db := libs.DB

	if db.Where("id = ?", id).First(&tag).RecordNotFound() {
		return Tags{}, errors.New("标签未找到")
	}
	return tag, nil
}

func (this *Tags) ListAll() []Tags {
	var data = []Tags{}
	db := libs.DB

	err := db.Order("sort desc").Find(&data).Error
	if err != nil {
		log.Fatalln(err)
	}

	return data
}

func (this *Tags) TagsAdd(postValues map[string][]string) error {
	var tag Tags
	db := libs.DB

	if err := libs.FormDecode(&tag, postValues); err != nil {
		return err
	}
	if err := libs.Validate(tag); err != nil {
		return err
	}
	if !db.Where("name = ? ", tag.Name).First(&Tags{}).RecordNotFound() {
		return errors.New("该名称已经存在")
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

func (this *Tags) TagsUpdate(postValues map[string][]string) error {
	var tag Tags
	db := libs.DB

	if err := libs.FormDecode(&tag, postValues); err != nil {
		return err
	}
	if err := libs.Validate(tag); err != nil {
		return err
	}
	// log.Println(tag, postValues)
	if !db.Where("name = ? and id != ?", tag.Name, tag.ID).Find(&Tags{}).RecordNotFound() {
		return errors.New("该名称已经存在")
	}
	if db.Where("id = ? ", tag.ID).Find(&Tags{}).RecordNotFound() {
		return errors.New("未查询到标签id")
	}
	if err := db.Save(&tag).Error; err != nil {
		return err
	}
	return nil
}

func (this *Tags) TagDel(id uint) error {
	var tag Tags

	db := libs.DB

	if err := db.Where("id = ?", id).Delete(&tag).Error; err != nil {
		return err
	}

	return nil
}

func (this *Tags) TagsMoreInfo(ids string) ([]Tags, error) {
	var data = []Tags{}
	db := libs.DB
	if db.Where("id in (?)", strings.Split(ids, ",")).Find(&data).RecordNotFound() {
		return []Tags{}, errors.New("未找到标签")
	}

	return data, nil
}

// func (this *Category) CategoryDel(id uint) error {
// 	var category Category
// 	db := libs.DB

// 	if !db.Where("parent_id = ?", id).Find(&category).RecordNotFound() {
// 		return errors.New("该分类下存在子级分类，请先删除子级分类")
// 	}
// 	if err := db.Where("id = ?", id).Delete(&category).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (this *Category) GetTree(list []Category, pid int, level int) []Category {
// 	for _, v := range list {
// 		if v.ParentId == pid {
// 			v.Level = level
// 			v.Name = strings.Repeat("————", v.Level) + v.Name
// 			ListTree = append(ListTree, v)
// 			/*if len(list) == 0 {
// 			      list = []Category{}
// 			  } else {
// 			      list = append(list[:index], list[index+1:]...)
// 			  }*/
// 			this.GetTree(list, int(v.Model.ID), v.Level+1)
// 		}
// 	}
// 	return ListTree
// }
