package middleware

import (
	"github.com/kataras/iris/context"
	"iris/commons"
	"log"
)

func SessionLoginAuth(Ctx context.Context) {
	auth := commons.SessManager.Start(Ctx).Get("admin_user")
	log.Println(auth)
	if auth == nil {
		Ctx.Redirect("/login")
		return
	}
	Ctx.Next()
}
