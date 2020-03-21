package middleware

import (
	"github.com/kataras/iris/context"
	config "github.com/spf13/viper"
	"iris/commons"
	"iris/model"
	"time"
)

// var session *sessions.Session

func SessionLoginAuth(Ctx context.Context) {
	auth := commons.SessManager.Start(Ctx).Get("admin_user")
	// log.Println(auth)
	if auth != nil {
		var userModel model.Admin
		user, _ := auth.(map[string]interface{})
		admin, _ := user["id"].(uint)
		password, _ := user["password"].(string)
		timeSession, _ := user["time"].(int64)
		timeNow := time.Now().Unix()
		timeDifference := timeNow - timeSession
		var timeOld int64 = config.GetInt64("site.SessionExpires") * 3600
		// log.Println(timeDifference < timeOld && userModel.CheckPassword(int(admin), password))
		if timeDifference < timeOld && userModel.CheckPassword(int(admin), password) {
			user["time"] = timeNow
			commons.SessManager.Start(Ctx).Set("admin_user", user)
			Ctx.Next()
			return
		}
		commons.SessManager.Start(Ctx).Delete("admin_user")
		Ctx.Redirect("/login")
		return
	} else {
		Ctx.Redirect("/login")
		return
	}
}
