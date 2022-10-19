package v1

import (
	"election/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 添加
type CandidateCreateReq struct {
	g.Meta       `path:"/candidate" method:"post" tags:"CandidateService" summary:"candidate create"`
	Name         string `v:"required" json:"name" description:"候选人名称"`
	Introduction string `v:"required" json:"introduction" description:"候选人简介"`
	Image        string `v:"required" json:"image" description:"候选人头像url"`
}
type CandidateCreateRes struct{}

// 更新
type CandidateUpdateReq struct {
	g.Meta       `path:"/candidate/:CandidateId" method:"put" tags:"CandidateService" summary:"candidate update"`
	Name         string `v:"required" json:"name" description:"候选人名称"`
	Introduction string `v:"required" json:"introduction" description:"候选人简介"`
	Image        string `v:"required" json:"image" description:"候选人头像url"`
	CandidateId  int    `v:"required" json:"candidateId" description:"候选人id"`
}
type CandidateUpdateRes struct{}

// 修改状态
type CandidateChangeStatusReq struct {
	g.Meta      `path:"/candidate/:CandidateId/status" method:"put" tags:"CandidateService" summary:"candidate changeStatus"`
	Status      int `v:"required" json:"status" description:"候选人状态 1-可用 0-不可用"`
	CandidateId int `v:"required" json:"name" description:"候选人Id"`
}
type CandidateChangeStatusRes struct{}

// 修改状态
type CandidateGetReq struct {
	g.Meta `path:"/candidate" method:"get" tags:"CandidateService" summary:"candidate get"`
	Page   int `d:"1"  v:"min:0#分页号码错误" json:"page" description:"页数"` // 分页号码
	Size   int `d:"10" v:"max:50#分页数量最大50条" json:"size" description:"每页条数"`
}
type CandidateGetRes struct {
	*model.CandidateGetOut
}
