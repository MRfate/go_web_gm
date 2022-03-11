package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Web struct {
}

func (web *Web) GetByte (url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("A error occurred!")
		return nil
	}
	defer res.Body.Close()
	// 读取获取的[]byte数据
	data, _ := ioutil.ReadAll(res.Body)
	return data
}