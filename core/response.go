package core

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Response 响应协议
type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	ReqId interface{} `json:"reqId"`
	Time  int64       `json:"t"`
}

// ErrorCode 错误码协议
type ErrorCode struct {
	Code int
	Msg  string
}

// ResSuccess 返回成功
func ResSuccess(c *gin.Context, data interface{}) {
	reqId, _ := c.Get("reqId")
	res := Response{
		Code:  0,
		Msg:   "OK",
		Data:  data,
		Time:  time.Now().Unix(),
		ReqId: reqId,
	}

	c.JSON(200, res)
	return
}

func ResSuccessWithMsg(c *gin.Context, data interface{}, msg string) {
	reqId, _ := c.Get("reqId")
	res := Response{
		Code:  0,
		Msg:   msg,
		Data:  data,
		Time:  time.Now().Unix(),
		ReqId: reqId,
	}

	c.JSON(200, res)
	return
}

// ResError 返回逻辑错误
func ResError(c *gin.Context, errorCode ErrorCode, subMsg string) {
	msg := errorCode.Msg
	if len(subMsg) > 0 {
		msg = msg + "(" + subMsg + ")"
	}

	reqId, _ := c.Get("reqId")

	res := Response{
		Code:  errorCode.Code,
		Msg:   msg,
		Data:  "",
		ReqId: reqId,
		Time:  time.Now().Unix(),
	}
	c.JSON(200, res)
	return
}

// ResError 返回逻辑错误
func ResErrorWithData(c *gin.Context, errorCode ErrorCode, subMsg string, data interface{}) {
	msg := errorCode.Msg
	if len(subMsg) == 0 {
		msg = msg + "(" + subMsg + ")"
	}

	reqId, _ := c.Get("reqId")

	res := Response{
		Code:  errorCode.Code,
		Msg:   msg,
		Data:  data,
		ReqId: reqId,
		Time:  time.Now().Unix(),
	}
	c.JSON(200, res)
	return
}
