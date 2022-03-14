package services

import (
	"github.com/kataras/iris/v12"
	"testIris2/utils"
)

type GMService struct {
}

func (gms *GMService) GetResize (scale int, url string, ext string) []byte {
	gm := utils.GMUtil{}
	web := utils.Web{}

	_byte := web.GetByte(url)
	img := gm.GetImageByByte(_byte)

	newImg := gm.ResizeImage(img, scale)
	return gm.ImageToByte(newImg, ext)
}

func (gms *GMService) ResizeByWidth(ctx iris.Context, url string, width int, ext string) ([]byte, bool) {
	gm := utils.GMUtil{}
	web := utils.Web{}
	_byte := web.GetByte(url)
	img := gm.GetImageByByte(_byte)
	size := gm.GetSize(img)
	if width > size.W {
		ctx.StatusCode(403)
		return nil, true
	}

	newImg := gm.ResizeImageByWidth(img, width)
	theByte := gm.ImageToByte(newImg, ext)
	return theByte, false
}


func (gms *GMService) ResizeImgByScaleInLocal(ctx iris.Context, url string, scale int) (string, interface{}, bool) {
	fileName := utils.GetLocalName(url)
	isExist, _ := utils.PathExists(fileName)
	if !isExist {
		ctx.StatusCode(404)
		return "", "文件不存在", true
	}

	// 获取缩略图的本地位置
	scaleImageName := utils.GetScaleImageName(fileName, scale)
	isExist, _ = utils.PathExists(scaleImageName)
	// 本地存在就直接重定向
	if isExist {
		url = utils.GetUrl(scaleImageName)
		ctx.Redirect(url, 302)
		return "", nil, true
	}

	gm := utils.GMUtil{}
	img := gm.GetImage(fileName)
	newImg := gm.ResizeImage(img, scale)
	gm.SaveImage(newImg, scaleImageName)
	return scaleImageName, nil, false
}


func (gms *GMService) ResizeImgByWidthInLocal(ctx iris.Context, url string, width int) (string, interface{}, bool) {
	fileName := utils.GetLocalName(url)
	isExist, _ := utils.PathExists(fileName)
	if !isExist {
		ctx.StatusCode(404)
		return "", "文件不存在", true
	}

	// 获取缩略图的本地位置
	widthImageName := utils.GetWidthImageName(fileName, width)
	isExist, _ = utils.PathExists(widthImageName)
	// 本地存在就直接重定向
	if isExist {
		url = utils.GetUrl(widthImageName)
		ctx.Redirect(url, 302)
		return "", nil, true
	}

	gm := utils.GMUtil{}
	img := gm.GetImage(fileName)
	newImg := gm.ResizeImageByWidth(img, width)
	gm.SaveImage(newImg, widthImageName)
	return widthImageName, nil, false
}
