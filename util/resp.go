package util

import (
	"encoding/json"
	"log"
)

// RespMsg 响应体结构
type RespMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// NewRespMsg 新建响应体
func NewRespMsg(code int, msg string, data interface{}) *RespMsg {
	return &RespMsg{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// JSONBytes RespMsg -> JSON Bytes
func (rm *RespMsg) JSONBytes() []byte {
	data, err := json.Marshal(rm)
	if err != nil {
		log.Println(err)
	}
	return data
}

// JSONString RespMsg -> JSON String
func (rm *RespMsg) JSONString() string {
	data := rm.JSONBytes()
	return string(data)
}
