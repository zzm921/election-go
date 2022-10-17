package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)
//创建选举
type ElectionCreateReq struct {
	g.Meta       `path:"/election" method:"post" tags:"ElectionService" summary:"election create"`
	Title        string `v:"required"`
	Introduction string `v:"required"`
	Candidates   []int  `v:"required"`
}
type ElectionCreateRes struct{}

//更新
type ElectionUpdateReq struct {
	g.Meta       `path:"/election/:electionId" method:"put" tags:"ElectionService" summary:"election create"`
	Title        string `v:"required"`
	Introduction string `v:"required"`
	Candidates   []int  `v:"required"`
}
type ElectionUpdateRes struct{}

//修改选举状态
type ElectionChangeStatusReq struct {
	g.Meta `path:"/election/:electionId/status" method:"put" tags:"ElectionService" summary:"election create"`
	Status int `v:"required"`
}
type ElectionChangeStatusRes struct{}

//获取选举
type ElectionGetReq struct {
	g.Meta `path:"/election" method:"get" tags:"ElectionService" summary:"election create"`
}
type ElectionGetRes struct{}


//获取选举
type ElectionVoteDetailGetReq struct {
	g.Meta `path:"/electionVoteDetail" method:"get" tags:"ElectionService" summary:"election create"`
	CandidateId   int  `v:"required"`
	ElectionId   int  `v:"required"`
}
type ElectionVoteDetailGetRes struct{}
