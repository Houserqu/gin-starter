package core

var (
	// 通用错误 < 1000
	ErrParam    = ErrorCode{Code: 400, Msg: "参数错误"}
	ErrNotFound = ErrorCode{Code: 404, Msg: "查找失败"}
	ErrDB       = ErrorCode{Code: 600, Msg: "数据库错误"}
	// 具体业务错误 > 1000
	ErrDeleteFail = ErrorCode{Code: 1000, Msg: "删除失败"}
	ErrCreateFail = ErrorCode{Code: 1001, Msg: "创建失败"}
	ErrUpdateFail = ErrorCode{Code: 1002, Msg: "更新失败"}
)
