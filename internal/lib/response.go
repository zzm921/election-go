package response

import (
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	OkCode                 = http.StatusOK
	OkMsg                  = "请求成功"
	ErrorCode              = http.StatusInternalServerError
	ErrorMsg               = "请求失败"
	IncorrectSignatureCode = 1005 // 数据已存在
	IncorrectSignatureMsg  = "无效签名"
	DataExistCode          = 1001 // 数据已存在
	DataExistMsg           = "数据已存在"
	ParamValidErrCode      = 1002      // 参数校验失败
	ParamValidErrMsg       = "参数校验失败"  // 参数校验失败
	DataNoExistCode        = 1003      // 帐号不存在
	DataNoExistMsg         = "数据不存在"   //
	AccountValidErrCode    = 1004      // 帐号或密码错误
	AccountValidErrMsg     = "帐号或密码错误" // 帐号或密码错误

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
