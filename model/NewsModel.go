package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	config "github.com/spf13/viper"
	"iris/libs"
	"iris/libs/logging"

	// "log"
	"math"
	"strings"
)

type News struct {
	gorm.Model
	Category_id  string `gorm:"type:varchar(100);DEFAULT '';"validate:"required"`
	Title        string `gorm:"type:varchar(250); NOT NULL; DEFAULT '';"validate:"required"`
	Descript     string `gorm:"type:varchar(500); NOT NULL; DEFAULT '0';"validate:"required"`
	Content      string `gorm:"type:text;NOT NULL; DEFAULT '0';"validate:"required,html"`
	Tags_id      string `gorm:"type:varchar(100);DEFAULT '';"`
	Sort         int    `gorm:"type:int(11); NOT NULL; DEFAULT '0';"validate:"number,min=0"`
	CategoryName string `gorm:"-"`
	TagsName     string `gorm:"-"`
}

func (this *News) List(page int) ([]News, int, int) {
	var data = []News{}
	var totalCount int
	db := libs.DB

	limit := config.GetInt("pagination.FrontendPageSize")
	offset := (page - 1) * limit
	db.Find(&data).Count(&totalCount)
	db.Offset(offset).Limit(limit).Order("id desc").Find(&data)
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	return data, totalCount, totalPages
}


func (this *News) ListBackend(page int) ([]News, int, int) {
	var data = []News{}
	var totalCount int
	db := libs.DB

	limit := config.GetInt("pagination.PageSize")
	offset := (page - 1) * limit
	db.Find(&data).Count(&totalCount)
	db.Offset(offset).Limit(limit).Order("id desc").Find(&data)
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	return data, totalCount, totalPages
}

func (this *News) NewsInfo(id uint) (News, error) {
	var news News
	db := libs.DB

	if db.Where("id = ?", id).First(&news).RecordNotFound() {
		return News{}, errors.New("内容未找到")
	}
	return news, nil
}

func (this *News) NewsAdd(postValues map[string][]string) error {
	// log.Println(postValues)
	var news News
	db := libs.DB

	if err := libs.FormDecode(&news, postValues); err != nil {
		libs.LogError.Println(err)
		return err
	}
	if err := libs.Validate(news); err != nil {
		libs.LogError.Println(err)
		return err
	}
	if !db.Where("title = ? ", news.Title).First(&News{}).RecordNotFound() {
		libs.LogInfo.Println("该标题已经存在")
		return errors.New("该标题已经存在")
	}

	news.Category_id = strings.Join(postValues["Category_id"], ",")
	news.Tags_id = strings.Join(postValues["tag_id"], ",")
	if err := db.Create(&news).Error; err != nil {
		libs.LogError.Println(err)
		return err
	}
	return nil
}

func (this *News) NewsUpdate(postValues map[string][]string) error {
	var news News
	db := libs.DB

	if err := libs.FormDecode(&news, postValues); err != nil {
		libs.LogError.Println(err)
		return err
	}
	if err := libs.Validate(news); err != nil {
		libs.LogError.Println(err)
		return err
	}
	if !db.Where("title = ? and id != ?", news.Title, news.ID).First(&News{}).RecordNotFound() {
		return errors.New("该标题已经存在")
	}
	if db.Where("id = ? ", news.ID).First(&News{}).RecordNotFound() {
		return errors.New("未查询到内容id")
	}

	news.Category_id = strings.Join(postValues["Category_id"], ",")
	news.Tags_id = strings.Join(postValues["tag_id"], ",")
	if err := db.Save(&news).Error; err != nil {
		libs.LogError.Println(err)
		return err
	}
	return nil
}

func (this *News) NewsDel(id uint) error {
	var news News
	db := libs.DB

	if err := db.Where("id = ?", id).Delete(&news).Error; err != nil {
		libs.LogError.Println(err)
		return err
	}
	return nil
}

func (this *News) NewsNewest() []News {
	var data = []News{}

	db := libs.DB

	err := db.Order("id desc").Limit(3).Find(&data).Error
	if err != nil {
		libs.LogError.Println(err)
	}

	return data

}

func (this *News) NewsSearch(search map[string]string) []News {
	var data = []News{}

	db := libs.DB

	for key, value := range search {
		db = db.Where("`news`.`"+key+"` like '%"+value+"%'")
	}
	err := db.Limit(config.Get("pagination.FrontendPageSize")).Find(&data).Error

	if err != nil {
		logging.Info("errors: ", err)

		return []News{}
	}

	return data
}
