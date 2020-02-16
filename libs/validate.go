package libs

import (
	"errors"
	"github.com/go-playground/form"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	// admin    *model.Admin
)

func Validate(s interface{}) error {
	zh := zh.New()
	trans, _ := ut.New(zh, zh).GetTranslator("zh")
	validate = validator.New()
	zh_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(s)
	errStr := ""
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			errStr += e.Translate(trans) + "\n"
		}
		return errors.New(errStr)
	}
	return nil
}

func FormDecode(v interface{}, postValues map[string][]string) error {
	var decoder *form.Decoder
	decoder = form.NewDecoder()
	err := decoder.Decode(&v, postValues)
	if err != nil {
		return err
	}
	return nil
}

//func CheckPassword(original map[string]interface {}) bool {
//	dbData := model.Admin{}
//	admin, err := dbData.AdminInfo(uint(original["id"]))
//	if err != nil {
//		log.Println("err: ", err)
//		return false
//	}
//	if admin.Password == original["password"] {
//		return true
//	}
//	return false
//}