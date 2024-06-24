package common

const OK int64 = 0
const SERVER_COMMON_ERROR int64 = 10001

type Errno struct {
	Code    int64
	Message string
}

// 用户模块
var (
	EMAIL_REGISTERED_ERROR = &Errno{Code: SERVER_COMMON_ERROR, Message: "邮箱已注册"}
	INSERT_USER_ERROR      = &Errno{Code: SERVER_COMMON_ERROR, Message: "用户注册失败"}
)
