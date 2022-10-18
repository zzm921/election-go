package v1

import (
	"election/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AccountSignUpReq struct {
	g.Meta   `path:"/account/login" method:"post" tags:"AccountService" summary:"account login"`
	Username string `v:"required"`
	Password string `v:"required"`
}
type AccountSignUpRes struct {
	*model.AccountLoginOut
}
