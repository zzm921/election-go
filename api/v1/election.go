package v1

import (
	"election/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 创建选举
type ElectionCreateReq struct {
	g.Meta       `path:"/election" method:"post" tags:"ElectionService" summary:"election create"`
	Title        string `v:"required"`
	Introduction string `v:"required"`
	Candidates   []int  `v:"required"`
}
type ElectionCreateRes struct{}

// 修改选举状态
type ElectionChangeStatusReq struct {
	g.Meta     `path:"/election/:ElectionId/status" method:"put" tags:"ElectionService" summary:"election create"`
	Status     int `v:"required"`
	ElectionId int `v:"required"`
}
type ElectionChangeStatusRes struct {
}

// 获取选举
type ElectionGetReq struct {
	g.Meta `path:"/election" method:"get" tags:"ElectionService" summary:"election create"`
	Page   int `d:"1"  v:"min:0#分页号码错误"` // 分页号码
	Size   int `d:"10" v:"max:50#分页数量最大50条"`
}
type ElectionGetRes struct {
	Count int
	List  []*model.ElectionGetOutListObject
}

// 获取选举
type ElectionVoteDetailGetReq struct {
	g.Meta      `path:"/electionVoteDetail" method:"get" tags:"ElectionService" summary:"election create"`
	CandidateId int `v:"required"`
	ElectionId  int `v:"required"`
}
type ElectionVoteDetailGetRes struct{}
