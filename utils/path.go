package utils

import (
	"os"
	"path"
	"strconv"
	"strings"
	"testIris2/config"
)
// GetLocalName 通过参数返回本地的文件路径
func GetLocalName(url string) string {
	return strings.Replace(url, config.Config.Prefix, config.Config.LocalPath, 1)
}

// GetUrl 通过文件路径返回对应的相对url
func GetUrl(fileName string) string {
	return strings.Replace(fileName, config.Config.LocalPath, config.Config.Prefix, 1)
}

// PathExists 判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetScaleImageName 通过一点的规则返回对应比例缩略图的文件位置
func GetScaleImageName(fileName string, scale int) string {
	ext := path.Ext(fileName)
	filePrefix := fileName[0:(len(fileName) - len(ext))]
	return filePrefix + "_scale_" + strconv.Itoa(scale) + ext
}

// GetWidthImageName 通过一点的规则返回对应比例缩略图的文件位置
func GetWidthImageName(fileName string, width int) string {
	ext := path.Ext(fileName)
	filePrefix := fileName[0:(len(fileName) - len(ext))]
	return filePrefix + "_width_" + strconv.Itoa(width) + ext
}