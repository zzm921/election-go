package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ElectionCreateReq struct {
	g.Meta       `path:"/election" method:"post" tags:"ElectionService" summary:"election create"`
	Title        string `v:"required"`
	Introduction string `v:"required"`
	Candidates   []int  `v:"required"`
}
type ElectionCreateRes struct{}

type ElectionUpdateReq struct {
	g.Meta       `path:"/election/:electionId" method:"put" tags:"ElectionService" summary:"election create"`
	Title        string `v:"required"`
	Introduction string `v:"required"`
	Candidates   []int  `v:"required"`
}
type ElectionUpdateRes struct{}

type ElectionChangeStatusReq struct {
	g.Meta `path:"/election/:electionId/status" method:"put" tags:"ElectionService" summary:"election create"`
	Status int `v:"required"`
}
type ElectionChangeStatusRes struct{}

type ElectionGetReq struct {
	g.Meta `path:"/election" method:"post" tags:"ElectionService" summary:"election create"`
}
type ElectionGetRes struct{}
