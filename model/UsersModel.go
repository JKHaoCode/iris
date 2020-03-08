package model

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	config "github.com/spf13/viper"
	"iris/libs"
	"iris/libs/logging"
	"log"
	"math"
	"sync"
)

var mu sync.Mutex

type Admin struct {
	gorm.Model
	//ID       uint    `gorm:"primary_key;auto_increment"`
	Account  string `gorm:"type:varchar(20);not null;index:username"`
	Password string `gorm:"type:char(32);not null;"`
	Nickname string `gorm:"type:char(100);DEFAULT '';"`
	Descript string `gorm:"type:varchar(255);DEFAULT '';"`
	Email    string `gorm:"type:varchar(100);DEFAULT '';"`
	Headico  string `gorm:"type:varchar(200);DEFAULT '';"`
	Online   int    `gorm:"type:bool;DEFAULT 1"`
}

func (this *Admin) List(page int) ([]Admin, int, int) {
	var data = []Admin{}
	var totalCount int
	limit := config.GetInt("pagination.PageSize")
	offset := (page - 1) * limit
	db := libs.DB.Table("users")
	db.Find(&data).Count(&totalCount)
	err := db.Offset(offset).Limit(limit).Order("id desc").Find(&data).Error
	if err != nil {
		log.Fatalln(err)
	}
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	return data, totalCount, totalPages
}

func (this *Admin) AdminLogin(account string, password string) (Admin, error) {
	if account == "" || password == "" {
		return Admin{}, errors.New("帐号或密码不能为空")
	}
	db := libs.DB.Table("users")
	var admin Admin
	has := md5.Sum([]byte(password))
	md5_password := fmt.Sprintf("%x", has) //将[]byte转成16进制
	if err := db.Model("Admin").Where("account = ?", account).Where("password = ?", md5_password).First(&admin).RecordNotFound(); err {
		logging.Info("登录错误： 没有查到", err)
		return Admin{}, errors.New("帐号或密码错误")
	}
	return admin, nil
}

func (this *Admin) AdminInfo(id uint) (Admin, error) {
	var admin Admin
	db := libs.DB.Table("users")
	if db.Where("id = ?", id).First(&admin).RecordNotFound() {
		return Admin{}, errors.New("用户未找到")
	}
	return admin, nil
}

func (this *Admin) AdminPasswodUpdate(admin_id uint, password, Repassword string) error {
	if password == "" || Repassword == "" {
		return errors.New("密码不能为空")
	}
	if password != Repassword {
		return errors.New("密码不一致")
	}

	db := libs.DB.Table("users")
	var admin Admin

	if db.Where("id = ? ", admin_id).First(&admin).RecordNotFound() {
		return errors.New("未查询到用户id")
	}

	has := md5.Sum([]byte(password))
	md5_password := fmt.Sprintf("%x", has) //将[]byte转成16进制

	if err := db.Model(&admin).Update("password", md5_password).Error; err != nil {
		return errors.New("密码修改失败")
	}

	return nil
}

func (this *Admin) AdminUpdate(postValues map[string][]string, admin_id uint, filePath string) error {
	var admin Admin
	db := libs.DB.Table("users")
	if db.Where("id = ? ", admin_id).First(&admin).RecordNotFound() {
		return errors.New("未查询到用户id")
	}

	admin.Nickname = postValues["nickname"][0]
	admin.Descript = postValues["descript"][0]
	admin.Email = postValues["email"][0]
	if filePath != "" {
		admin.Headico = filePath
	}
	if err := db.Save(&admin).Error; err != nil {
		return errors.New("修改失败")
	}

	return nil
}

func (this *Admin) AddUpdate(postValues map[string][]string, filePath string) error {
	var admin Admin

	if postValues["password"][0] == "" || postValues["Repassword"][0] == "" {
		return errors.New("密码不能为空")
	}
	if postValues["password"][0] != postValues["Repassword"][0] {
		return errors.New("密码不一致")
	}
	delete(postValues, "Repassword")
	has := md5.Sum([]byte(postValues["password"][0]))
	postValues["password"][0] = fmt.Sprintf("%x", has) //将[]byte转成16进制

	db := libs.DB.Table("users")

	if !db.Where("account = ? ", postValues["account"][0]).First(&admin).RecordNotFound() {
		return errors.New("该账户已经存在")
	}

	admin.Account = postValues["account"][0]
	admin.Password = postValues["password"][0]
	admin.Nickname = postValues["nickname"][0]
	admin.Descript = postValues["descript"][0]
	admin.Email = postValues["email"][0]
	if filePath != "" {
		admin.Headico = filePath
	}
	if err := db.Create(&admin).Error; err != nil {
		return errors.New("新增失败")
	}

	return nil
}

func (this *Admin) AdminDel(id uint) error {
	AdminId := uint(config.GetInt("site.AdminId"))
	if id == AdminId {
		return errors.New("系统管理员不允许删除")
	}
	var admin Admin
	db := libs.DB.Table("users")
	if err := db.Where("id = ?", id).Delete(&admin).Error; err != nil {
		return errors.New("删除失败")
	}
	return nil
}


func (this *Admin)CheckPassword(id int, password string) bool {
	// admin_user, _ := original.(map[string]interface{})
	// dbData := .Admin{}
	admin, err := this.AdminInfo(uint(id))
	// log.Println(admin)
	if err != nil {
		// log.Println("err: ", err)
		return false
	}
	if admin.Password == password {
		return true
	}
	return false
}

func (this *Admin)ChangeOnline(status int, id uint) bool {
	db := libs.DB.Table("users")
	var online int
	if status == 1 {
		online = 0
	} else {
		online = 1
	}
	err := db.Where("id = ?", id).Update("online", online).Error
	if err != nil {
		logging.Info("change online off: ", err)
		return false
	}

	return true
}