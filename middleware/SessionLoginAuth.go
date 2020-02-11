package middleware

import (
	"github.com/kataras/iris/context"
	"iris/commons"
)

func SessionLoginAuth(Ctx context.Context) {
	if auth := commons.SessManager.Start(Ctx).Get("admin_user"); auth == nil {
		Ctx.Redirect("/login")
		return
	}
	Ctx.Next()
}
