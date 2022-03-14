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

func (g *GMController) GetLocalWidth(ctx iris.Context) interface{} {
	query := ctx.Request().URL.Query()
	url := query.Get("url")
	width, err := strconv.Atoi(query.Get("width"))
	if err != nil {
		fmt.Println(err)
		ctx.StatusCode(404)
		return nil
	}
	if url == "" || width <= 0 || !strings.HasPrefix(url, config.Config.Prefix) {
		ctx.StatusCode(404)
		return nil
	}
	theGMService := services.GMService{}
	scaleImageName, i, done := theGMService.ResizeImgByWidthInLocal(ctx, url, width)
	if done {
		return i
	}
	ctx.Redirect(utils.GetUrl(scaleImageName), 302)
	return nil
}


func (g *GMController) GetLocalScale(ctx iris.Context) interface{} {
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
	theGMService := services.GMService{}
	scaleImageName, i, done := theGMService.ResizeImgByScaleInLocal(ctx, url, scale)
	if done {
		return i
	}
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

func (g *GMController) GetWidth(ctx iris.Context) {
	query := ctx.Request().URL.Query()
	url := query.Get("url")
	width, err := strconv.Atoi(query.Get("width"))
	if err != nil {
		fmt.Println(err)
		ctx.StatusCode(404)
		return
	}
	if url == "" || width <= 0 {
		ctx.StatusCode(404)
		return
	}
	ext := strings.ToLower(path.Ext(url))

	theGMService := services.GMService{}
	theByte, done := theGMService.ResizeByWidth(ctx, url, width, ext)
	if done {
		return
	}
	_, err = ctx.Write(theByte)
	if err != nil {
		fmt.Println(err)
		ctx.StatusCode(500)
		return
	}
	ctx.Header("Content-Type", "image/"+ext[1:len(ext)])
	ctx.ResponseWriter()
}
