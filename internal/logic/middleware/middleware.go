package middleware

import (
	response "election/internal/lib"
	"election/internal/model"
	"election/internal/service"
	"encoding/json"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) AccountAuth(r *ghttp.Request) {
	accountToken := r.Cookie.Get("accountToken").String()
	if accountToken == "" {
		//未登录返回
		r.Response.WriteJson(g.Map{
			"code":    gcode.CodeNotAuthorized.Code(),
			"message": gcode.CodeNotAuthorized.Message(),
			"data":    nil,
		})

		r.ExitAll()
	}
	//通过token获取到登录信息
	accountDataStr, _ := g.Redis().Do(r.Context(), "get", accountToken)
	if accountDataStr.String() == "" {
		//未登录返回
		r.Response.WriteJson(g.Map{
			"code":    response.AuthorizedFailCode,
			"message": response.AuthorizedFailMsg,
			"data":    nil,
		})
	}
	//转为struct
	var accountData *model.AccountLoginOut
	json.Unmarshal([]byte(accountDataStr.String()), &accountData)
	r.SetCtxVar("account", accountData)
	r.Middleware.Next()
}
