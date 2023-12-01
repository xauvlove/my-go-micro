package utils

const (
	SYSTEM_ERROR = "1"
	RPC_ERROR    = "2"
	CHECK_FAILD  = "3"
	SUCCESS      = "4"
)

var recodeText = map[string]string{
	SYSTEM_ERROR: "系统错误",
	RPC_ERROR:    "远程服务异常",
	CHECK_FAILD:  "验证失败",
	SUCCESS:      "成功",
}

func RecodeText(code string) string {
	s, ok := recodeText[code]
	if ok {
		return s
	}
	return "未知异常"
}
