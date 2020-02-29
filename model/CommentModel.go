package model

import (
	"github.com/jinzhu/gorm"
	config "github.com/spf13/viper"
	"iris/libs"
	"iris/libs/logging"
	"log"
	"math"
)

type Comment struct {
	gorm.Model
	UserId uint `gorm:"type:int(10);DEFAULT 0;"`
	ArticleId uint `gorm:"type:int(10); NOT NULL; DEFAULT 0;"validate:"required"`
	CommentLikeCount uint `gorm:"type:bigint(20);DEFAULT 0;"`
	CommentContent string `gorm:"type:text;DEFAULT '';"validate:"required"`
	ArticleName string `gorm:"-"`
}

func (this *Comment) List(page int) ([]Comment, int, int){
	var data = []Comment{}
	var totalCount int
	limit := config.GetInt("pagination.PageSize")
	db := libs.DB

	offset := (page - 1) * limit
	db.Find(&data).Count(&totalCount)
	db.Offset(offset).Limit(limit).Order("id desc").Find(&data)
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return data, totalCount, totalPages
}

func (this *Comment) CommentAdd(postValues map[string][]string) error {
	var comment Comment
	db := libs.DB
	// log.Println(postValues)
	if err := libs.FormDecode(&comment, postValues); err != nil {
		libs.LogError.Println(err)
		log.Println("40", err)
		return err
	}
	if err := libs.Validate(comment); err != nil {
		log.Println("44", err)
		libs.LogError.Println(err)
		return err
	}

	if err := db.Create(&comment).Error; err != nil {
		log.Println("50:", err)
		libs.LogError.Println(err)
		return err
	}
	return nil
}

func (this *Comment) CommentSearch(search uint) []Comment {
	var data = []Comment{}
	db := libs.DB
	err := db.Where("article_id = ?", search).Order("id desc").Find(&data).Error
	if err != nil {
		logging.Info("errors: ", err)
		return []Comment{}
	}

	return data
}

func (this *Comment) ChangeCommentLike(total uint, id uint) bool {
	db := libs.DB.Table("comment")

	err := db.Where("id = ?", id).Update("comment_like_count", total + 1).Error

	if err != nil {
		logging.Info("change commentCount off: ", err)
		return false
	}
	return true
}