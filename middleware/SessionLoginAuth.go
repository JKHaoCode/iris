package middleware

import (
	"github.com/kataras/iris/context"
	"iris/commons"
	"iris/model"
	"log"
	"strconv"
)

func SessionLoginAuth(Ctx context.Context) {
	auth := commons.SessManager.Start(Ctx).Get("admin_user")
	log.Println("auth:", auth)

	// var checkedAuth map[string] string

	if auth != nil {
		var userModel model.Admin
		user, _ := auth.(map[string]interface{})
		log.Println("users: ", user)
		checkedUser := map[string] string{}
		admin, _ := user["id"].(uint)
		password, _ := user["password"].(string)
		userId := strconv.FormatUint(uint64(admin), 10)
		checkedUser["id"] = userId
		checkedUser["password"] = password
		// log.Println(checkedUser)
		log.Println("checked: ", checkedUser)
		if userModel.CheckPassword(checkedUser) {
			//Ctx.Redirect("/login")
			//return
			Ctx.Next()
		} else {
			Ctx.Redirect("/login")
		}
		// Ctx.Next()
	}
	// log.Println(libs.CheckPassword(auth))
	if auth == nil {
		Ctx.Redirect("/login")
		return
	}
	// Ctx.Next()
}
