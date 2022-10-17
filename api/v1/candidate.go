package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)
//创建选举
type CandidateCreateReq struct {
	g.Meta       `path:"/candidate" method:"post" tags:"CandidateService" summary:"candidate create"`
	Name        string `v:"required"`
	Introduction string `v:"required"`
	Image   string  `v:"required"`
}
type CandidateCreateRes struct{}

//更新
type CandidateUpdateReq struct {
	g.Meta       `path:"/candidate/:candidateId" method:"put" tags:"CandidateService" summary:"candidate create"`
	Name        string `v:"required"`
	Introduction string `v:"required"`
	Image   string  `v:"required"`
}
type CandidateUpdateRes struct{}

