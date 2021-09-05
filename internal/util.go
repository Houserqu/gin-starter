package internal

import "time"

// Response 响应协议
type Response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time int64       `json:"t"`
}

// ErrorCode 错误码协议
type ErrorCode struct {
	Code string
	Msg  string
}

// ResSuccess 返回成功
func ResSuccess(data interface{}, msg string) (status int, res Response) {
	status = 200
	if len(msg) == 0 {
		msg = "成功"
	}

	res = Response{
		Code: "",
		Msg:  msg,
		Data: data,
		Time: time.Now().Unix(),
	}
	return
}

// ResError 返回逻辑错误
func ResError(errorCode ErrorCode, data interface{}) (status int, res Response) {
	status = 200
	msg := errorCode.Msg
	res = Response{
		Code: errorCode.Code,
		Msg:  msg,
		Data: data,
		Time: time.Now().Unix(),
	}

	return
}

func ResParamError(data interface{}) (status int, res Response) {
	return ResError(ErrParam, data)
}

func ResNotFound(data interface{}) (status int, res Response) {
	return ResError(ErrNotFound, data)
}
