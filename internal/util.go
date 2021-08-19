package internal

import "time"

// Response 响应协议
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time int64       `json:"t"`
}

// ErrorCode 错误码协议
type ErrorCode struct {
	Code int
	Msg  string
}

// Success 返回成功
func Success(data interface{}, msg string) (status int, res Response) {
	status = 200
	res = Response{
		Code: 0,
		Msg:  msg,
		Data: data,
		Time: time.Now().Unix(),
	}
	return
}

// Error 返回逻辑错误
func Error(errorCode ErrorCode, other ...string) (status int, res Response) {
	status = 200

	msg := errorCode.Msg

	if len(other) > 0 {
		msg += "(" + other[0] + ")"
	}

	res = Response{
		Code: errorCode.Code,
		Msg:  msg,
		Time: time.Now().Unix(),
	}

	return
}
