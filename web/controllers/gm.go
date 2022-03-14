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
	gmService := services.GMService{}
	scaleImageName, i, done := gmService.ResizeImgByWidthInLocal(ctx, url, width)
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
	gmService := services.GMService{}
	scaleImageName, i, done := gmService.ResizeImgByScaleInLocal(ctx, url, scale)
	if done {
		return i
	}
	ctx.Redirect(utils.GetUrl(scaleImageName), 302)
	return nil
}

func (g *GMController) GetLocalSub(ctx iris.Context) {
	query := ctx.Request().URL.Query()
	url := query.Get("url")
	x, _ := strconv.Atoi(query.Get("x"))
	y, err := strconv.Atoi(query.Get("y"))
	w, err := strconv.Atoi(query.Get("w"))
	h, err := strconv.Atoi(query.Get("h"))
	if x < 0 || y < 0 || w < 1 || h < 1 {
		ctx.StatusCode(404)
		return
	}
	if err != nil {
		fmt.Println(err)
		ctx.StatusCode(404)
		return
	}
	fmt.Println(x, y, w, h, url)
	gmService := services.GMService{}
	subImgName, done := gmService.SubImage(ctx, url, x, y, w, h, err)
	if done {
		return
	}
	ctx.Redirect(utils.GetUrl(subImgName), 302)
	return
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
	gmService := services.GMService{}
	theByte := gmService.GetResize(scale, url, ext)

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

	gmService := services.GMService{}
	theByte, done := gmService.ResizeByWidth(ctx, url, width, ext)
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
