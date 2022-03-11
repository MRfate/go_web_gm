package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"path"
	"strconv"
	"strings"
	"testIris2/config"
	"testIris2/datamodels"
	"testIris2/services"
	"testIris2/utils"
)

type GMController struct {
}

func (g *GMController) Get() datamodels.GMRes {
	res := datamodels.GMRes{}
	res.Path = "hello"
	return res
}

func (g *GMController) GetImageScale(ctx iris.Context) interface{} {
	query := ctx.Request().URL.Query()
	url := query.Get("url")
	scale, err := strconv.Atoi(query.Get("scale"))
	if err != nil {
		fmt.Println(err)
		ctx.StatusCode(404)
		return nil
	}
	if url == "" || scale >= 100 || scale <= 1 || !strings.HasPrefix(url, config.Config.Prefix) {
		ctx.StatusCode(404)
		return nil
	}

	fileName := utils.GetLocalName(url)
	isExist, _ := utils.PathExists(fileName)
	if !isExist {
		ctx.StatusCode(404)
		return "文件不存在"
	}

	// 获取缩略图的本地位置
	scaleImageName := utils.GetScaleImageName(fileName, scale)
	isExist, _ = utils.PathExists(scaleImageName)
	// 本地存在就直接重定向
	if isExist {
		url = utils.GetUrl(scaleImageName)
		ctx.Redirect(url, 302)
		return nil
	}

	gm := utils.GMUtil{}
	img := gm.GetImage(fileName)
	newImg := gm.ResizeImage(img, scale)
	gm.SaveImage(newImg, scaleImageName)
	ctx.Redirect(utils.GetUrl(scaleImageName), 302)
	return nil
}

func (g *GMController) GetResize(ctx iris.Context) {
	query := ctx.Request().URL.Query()
	url := query.Get("url")
	scale, err := strconv.Atoi(query.Get("scale"))
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
	ctx.Header("Content-Type", "image/"+ext[1:len(ext)])
	ctx.ResponseWriter()
}
