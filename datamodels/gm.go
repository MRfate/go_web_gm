package datamodels

type GMRes struct {
	Path string `json:"path" bson:"path"`
	Num int64 `json:"num" bson:"num"`
}
