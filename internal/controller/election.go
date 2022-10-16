package controller

import (
	"context"
	"election/internal/model"
	"election/internal/service"

	v1 "election/api/v1"

	"github.com/gogf/gf/v2/frame/g"
)

var Electoin = cElection{}

type cElection struct{}

func (c *cElection) Create(ctx context.Context, req *v1.ElectionCreateReq) (res *v1.ElectionCreateRes, err error) {
	responseResult := service.Election().Create(ctx, model.ElectionCreateInput{
		Title:        req.Title,
		Introduction: req.Introduction,
		Candidates:   req.Candidates,
	})
	g.RequestFromCtx(ctx).Response.WriteJson(responseResult)
	return
}
