package lib

// Response 响应协议
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 返回成功内容方法
func Success(data interface{}, msg string) (status int, res Response) {
	status = 200
	res = Response{Code: 0, Msg: msg, Data: data}
	return
}
