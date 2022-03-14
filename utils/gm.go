package utils

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"strings"
)

type Size struct {
	W int `json:"w"`
	H int `json:"h"`
}

type GMUtil struct {
}

// SaveImage 保存图片到本地
func (gm *GMUtil) SaveImage(img image.Image, outPath string) {
	err := imaging.Save(img, outPath)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// ResizeImage 按比例生成缩略图
func (gm *GMUtil) ResizeImage(img image.Image, scale int) image.Image {
	if scale > 100 {
		return nil
	}
	size := gm.GetSize(img)
	w := size.W * scale / 100
	return gm.ResizeImageByWidth(img, w)
}


// ResizeImageByPath 按比例生成缩略图
func (gm *GMUtil) ResizeImageByPath(path string, scale int) image.Image {
	if scale > 100 {
		return nil
	}
	img := gm.GetImage(path)
	size := gm.GetSize(img)
	w := size.W * scale / 100
	return gm.ResizeImageByWidth(img, w)
}

// ResizeImageByWidth 按宽度生成缩略图
func (gm *GMUtil) ResizeImageByWidth (img image.Image, w int) image.Image {
	return imaging.Resize(img, w, 0, imaging.Lanczos)
}

// GetImage 通过路径获取图片对象
func (gm *GMUtil) GetImage(path string) image.Image {
	imgData, _ := ioutil.ReadFile(path)
	buf := bytes.NewBuffer(imgData)
	img, err := imaging.Decode(buf)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return img
}

// GetImageByByte 通过byte获取图片对象
func (gm *GMUtil) GetImageByByte (theByte []byte) image.Image {
	buf := bytes.NewBuffer(theByte)
	img, err := imaging.Decode(buf)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return img
}

func (gm *GMUtil) ImageToByte(img image.Image, ext string) []byte {
	ext = strings.ToLower(ext)
	buf := new(bytes.Buffer)
	if strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg") {
		err := jpeg.Encode(buf, img, &jpeg.Options{Quality: 80})
		if err != nil {
			return nil
		}
	} else if strings.EqualFold(ext, ".png") {
		err := png.Encode(buf, img)
		if err != nil {
			return nil
		}
	} else if strings.EqualFold(ext, ".gif") {
		err := gif.Encode(buf, img, &gif.Options{NumColors: 256})
		if err != nil {
			return nil
		}
	}
	return buf.Bytes()
}

// GetImageByBuf 通过路径获取图片对象
func (gm *GMUtil) GetImageByBuf(buf io.Reader) image.Image {
	img, err := imaging.Decode(buf)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return img
}

func (gm *GMUtil) GetSize(img image.Image) Size {
	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y
	return Size{
		H: height,
		W: width,
	}
}
