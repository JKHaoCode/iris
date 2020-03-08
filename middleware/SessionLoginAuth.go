package middleware

import (
	"github.com/kataras/iris/context"
	"iris/commons"
	"iris/model"
)

// var session *sessions.Session

func SessionLoginAuth(Ctx context.Context) {
	auth := commons.SessManager.Start(Ctx).Get("admin_user")

	if auth != nil {
		var userModel model.Admin
		user, _ := auth.(map[string]interface{})
		admin, _ := user["id"].(uint)
		password, _ := user["password"].(string)
		if userModel.CheckPassword(int(admin), password) {
			Ctx.Next()
		} else {
			Ctx.Redirect("/login")
			return
		}
	} else {
		Ctx.Redirect("/login")
		return
	}
}
