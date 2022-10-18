package v1

import (
	"election/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 添加
type CandidateCreateReq struct {
	g.Meta       `path:"/candidate" method:"post" tags:"CandidateService" summary:"candidate create"`
	Name         string `v:"required"`
	Introduction string `v:"required"`
	Image        string `v:"required"`
}
type CandidateCreateRes struct{}

// 更新
type CandidateUpdateReq struct {
	g.Meta       `path:"/candidate/:CandidateId" method:"put" tags:"CandidateService" summary:"candidate update"`
	Name         string `v:"required"`
	Introduction string `v:"required"`
	Image        string `v:"required"`
}
type CandidateUpdateRes struct{}

// 修改状态
type CandidateChangeStatusReq struct {
	g.Meta      `path:"/candidate/:CandidateId/status" method:"put" tags:"CandidateService" summary:"candidate changeStatus"`
	Status      int `v:"required"`
	CandidateId int `v:"required"`
}
type CandidateChangeStatusRes struct{}

// 修改状态
type CandidateGetReq struct {
	g.Meta `path:"/candidate" method:"get" tags:"CandidateService" summary:"candidate get"`
	Page   int `d:"1"  v:"min:0#分页号码错误"` // 分页号码
	Size   int `d:"10" v:"max:50#分页数量最大50条"`
}
type CandidateGetRes struct {
	*model.CandidateGetOut
}
