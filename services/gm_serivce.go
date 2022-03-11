package services

import (
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