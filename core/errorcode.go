package core

var (
	ErrParam      = ErrorCode{Code: "INVALID_PARAM", Msg: "参数错误"}
	ErrNotFound   = ErrorCode{Code: "NOT_FOUND", Msg: "查找失败"}
	ErrDeleteFail = ErrorCode{Code: "DELETE_FAIL", Msg: "删除失败"}
	ErrCreateFail = ErrorCode{Code: "CREATE_FAIL", Msg: "创建失败"}
	ErrUpdateFail = ErrorCode{Code: "UPDATE_FAIL", Msg: "更新失败"}
)
