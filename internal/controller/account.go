package controller

import (
	"context"
	"election/internal/model"
	"election/internal/service"
	"encoding/json"

	v1 "election/api/v1"

	"github.com/gogf/gf/v2/frame/g"
)

var Account = cAccount{}

type cAccount struct{}

func (c *cAccount) Login(ctx context.Context, req *v1.AccountSignUpReq) (res *v1.AccountSignUpRes, err error) {
	accountLoginOut, err := service.Account().Login(ctx, model.AccountLoginInput{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	accountJsonStr, _ := json.Marshal(&accountLoginOut)
	//将登陆信息存入redis中
	_, err = g.Redis().Do(ctx, "SETEX", "token", 86400, string(accountJsonStr))
	if err != nil {
		return nil, err
	}
	return &v1.AccountSignUpRes{AccountLoginOut: accountLoginOut}, nil
}
