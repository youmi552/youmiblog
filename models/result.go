package models

type Result struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Code int         `json:"code"`
}
