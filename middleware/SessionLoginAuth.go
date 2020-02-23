package middleware

import (
	"github.com/kataras/iris/context"
	"iris/commons"
	"iris/model"
)

func SessionLoginAuth(Ctx context.Context) {
	auth := commons.SessManager.Start(Ctx).Get("admin_user")
	// log.Println("auth:", auth)

	// var checkedAuth map[string] string

	if auth != nil {
		var userModel model.Admin
		user, _ := auth.(map[string]interface{})
		// log.Println("users: ", user)
		// checkedUser := map[string] string{}
		admin, _ := user["id"].(float64)
		password, _ := user["password"].(string)
		//log.Println("errs int", err)
		//fmt.Printf("type %T\n", user["id"])
		//fmt.Printf("type %T, %f, %d, %T\n", admin, admin, int(admin), int(admin))
		//log.Println("checked: ", user["id"], admin)
		if userModel.CheckPassword(int(admin), password) {
			Ctx.Next()
		} else {
			Ctx.Redirect("/login")
			return
		}
		// Ctx.Next()
	} else {
		Ctx.Redirect("/login")
		return
	}
	// log.Println(libs.CheckPassword(auth))
	//if auth == nil {
	//	Ctx.Redirect("/login")
	//	return
	//}
	// Ctx.Next()
}
