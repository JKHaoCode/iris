package counter

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/host"
	"time"
)

func Configurator(app *iris.Application) {
	counterValue := 0

	go func() {
		ticker := time.NewTicker(time.Second)

		for range ticker.C {
			counterValue++
		}

		app.ConfigureHost(func(h *host.Supervisor) { // <- 这里: 很重要
			h.RegisterOnShutdown(func() {
				ticker.Stop()
			})
		})
	}()

	app.Get("/counter", func(ctx iris.Context) {
		ctx.Writef("Counter value = %d", counterValue)
	})
}
