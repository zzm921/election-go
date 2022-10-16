package controller

import (
	"context"
	"election/internal/model"
	"election/internal/service"

	v1 "election/api/v1"

	"github.com/gogf/gf/v2/frame/g"
)

var Account = cAccount{}

type cAccount struct{}

func (c *cAccount) Login(ctx context.Context, req *v1.AccountSignUpReq) (res *v1.AccountSignUpRes, err error) {
	responseResult := service.Account().Login(ctx, model.AccountLoginInput{
		Username: req.Username,
		Password: req.Password,
	})
	g.RequestFromCtx(ctx).Response.WriteJson(responseResult)
	return
}
