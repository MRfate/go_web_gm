package main

import (
	"github.com/kataras/iris/v12"
	"testIris2/utils"
	"testIris2/web/controllers"
)
import "github.com/kataras/iris/v12/middleware/logger"
import "github.com/kataras/iris/v12/middleware/recover"
import "github.com/kataras/iris/v12/mvc"

func test(path string)  {
	gm := utils.GMUtil{}
	webUtil := utils.Web{}
	theByte := webUtil.GetByte(path)
	img := gm.GetImageByByte(theByte)
	img2 := gm.ResizeImage(img, 50)
	gm.SaveImage(img2, "./testFile/2.png")
}


func main() {
	app := iris.New()
	app.Use(logger.New())
	app.Use(recover.New())
	// test("https://img02.mockplus.cn/image/2022-1-29/c34deac0-80a7-11ec-9336-0356366d5597.png")
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
