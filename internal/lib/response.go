package lib

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	DataExistCode = 1001 // 数据已存在
	DataExistMsg  = "数据已存在"

	DataNoExistCode = 1002    // 帐号不存在
	DataNoExistMsg  = "数据不存在" //

	AuthorizedFailCode = 1003
	AuthorizedFailMsg  = "无效token"
)

type ResultRes struct {
	Code    int         `json:"code"` // response code, default: success-200 error-500 响应码
	Message string      `json:"msg"`  // response message 响应信息，若返回的是错误码，则此处对应相关错误信息
	Data    interface{} `json:"data"` // response result data 返回数据
}

// RestResult to write response, but not exit. 返回统一格式对象，但不退出。
func RestResult(r *ghttp.Request, code int, msg string, data interface{}) {
	r.Response.WriteJson(ResultRes{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
