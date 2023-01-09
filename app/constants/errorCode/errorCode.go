package errorCode

const (
	SUCCESS      = 200
	SERVER_ERROR = 500
	LOGIN_ERROR  = 401
	BAD_REQUEST  = 400
	AUTH_FAIL    = 402

	NOT_FOUND_USER = 4001
)

var MsgFlags = map[int]string{
	SUCCESS:      "请求成功",
	SERVER_ERROR: "服务器错误",
	LOGIN_ERROR:  "登录失败，请检查用户名或密码",
	BAD_REQUEST:  "请求失败",
	AUTH_FAIL:    "token 异常",

	NOT_FOUND_USER: "未找到用户",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[SERVER_ERROR]
}
