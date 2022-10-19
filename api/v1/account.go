package v1

import (
	"election/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AccountSignUpReq struct {
	g.Meta   `path:"/account/login" method:"post" tags:"AccountService" summary:"account login"`
	Username string `v:"required"  json:"username" description:"账号名"`
	Password string `v:"required"  json:"password" description:"密码"`
}
type AccountSignUpRes struct {
	*model.AccountLoginOut
}
