package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"path"
	"strconv"
	"strings"
	"testIris2/datamodels"
	"testIris2/services"
)


type GMController struct {
}

func (g *GMController) Get() datamodels.GMRes {
	res := datamodels.GMRes{}
	res.Path = "hello"
	return res
}


func (g *GMController) GetResize(ctx iris.Context) {
	a := ctx.Request().URL.Query()
	url := a.Get("url")
	scale, err := strconv.Atoi(a.Get("scale"))
	if err != nil {
		fmt.Println(err)
		ctx.StatusCode(404)
		return
	}
	if url == "" || scale >= 100 || scale <= 1 {
		ctx.StatusCode(404)
		return
	}
	ext := strings.ToLower(path.Ext(url))
	theGMService := services.GMService{}
	theByte := theGMService.GetResize(scale, url, ext)

	_, err = ctx.Write(theByte)
	if err != nil {
		fmt.Println(err)
		ctx.StatusCode(500)
		return
	}
	ctx.Header("Content-Type", "image/"+ext[1 : len(ext)])
	ctx.ResponseWriter()
}
