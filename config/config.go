package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ConfigModel struct {
	LocalPath string `json:"localPath"`
	Prefix    string `json:"prefix"`
}

var Config = &ConfigModel{}

const PATH string = "./config/config.json"

func loadConfig() *ConfigModel {
	data, err := ioutil.ReadFile(PATH)
	if err != nil {
		log.Fatal(err)
	}
	config := &ConfigModel{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func init() {
	Config = loadConfig()
}
