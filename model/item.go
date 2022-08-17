package model

type Item struct {
	Key  string      `json:"key"`
	Data interface{} `json:"data"`
}
