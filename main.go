package main

import (
	"github.com/kataras/iris/v12"
	"testIris2/config"
	"testIris2/web/controllers"
)
import "github.com/kataras/iris/v12/middleware/logger"
import "github.com/kataras/iris/v12/middleware/recover"
import "github.com/kataras/iris/v12/mvc"

func main() {
	app := iris.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Get("/", func(ctx iris.Context){
		_, err := ctx.JSON(iris.Map{
			"message": "hello world",
		})
		if err != nil {
			return 
		}
	})

	// 加载控制器
	mvc.Configure(app.Party("/gm"), gm)
	app.HandleDir(config.Config.Prefix, iris.Dir(config.Config.LocalPath))
	err := app.Run(iris.Addr(":8080"))
	if err != nil {
		return
	}
}

// 注意 mvc.Application, 不是 iris.Application.
func gm(app *mvc.Application) {
	//初始化控制器
	app.Handle(new(controllers.GMController))
}
