package v1

import (
	"election/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 创建选举
type ElectionCreateReq struct {
	g.Meta       `path:"/election" method:"post" tags:"ElectionService" summary:"election create"`
	Title        string `v:"required" json:"title" description:"选举标题"`
	Introduction string `v:"required" json:"introduction" description:"选举简介"`
	Candidates   []int  `v:"required" json:"candidates" description:"候选人id列表"`
}
type ElectionCreateRes struct{}

// 修改选举状态
type ElectionChangeStatusReq struct {
	g.Meta     `path:"/election/:ElectionId/status" method:"put" tags:"ElectionService" summary:"election changeStatus"`
	Status     int `v:"in:0,1,2" json:"Status" description:"选举状态 0-未开始 1-进行中 2-结束"`
	ElectionId int `v:"required" json:"electionId" description:"选举id"`
}
type ElectionChangeStatusRes struct {
}

// 获取选举候选人得票详情
type ElectionCandidateVoteGetReq struct {
	g.Meta      `path:"/election/:ElectionId/candidates/:CandidateId/vote" method:"get" tags:"ElectionService" summary:"election candidatesVoteGet"`
	ElectionId  int `v:"required" json:"electionId" description:"选举id"`
	CandidateId int `v:"required" json:"candidateId" description:"候选人id"`
	Page        int `d:"1"  v:"min:0#分页号码错误" json:"page" description:"页数"` // 分页号码
	Size        int `d:"10" v:"max:50#分页数量最大50条" json:"size" description:"每页条数"`
}
type ElectionCandidateVoteGetRes struct {
	*model.ElectionCandidateVoteGetOut
}

// 获取选举候选人信息
type ElectionCandidateGetReq struct {
	g.Meta     `path:"/election/:ElectionId/candidates" method:"get" tags:"ElectionService" summary:"election candidatesGet"`
	ElectionId int `v:"required" json:"electionId" description:"选举id"`
	Page       int `d:"1"  v:"min:0#分页号码错误" json:"page" description:"页数"` // 分页号码
	Size       int `d:"10" v:"max:50#分页数量最大50条" json:"size" description:"每页条数"`
}
type ElectionCandidateGetRes struct {
	*model.ElectionCandidateGetOut
}

// 获取选举
type ElectionGetReq struct {
	g.Meta `path:"/election" method:"get" tags:"ElectionService" summary:"election get"`
	Page   int `d:"1"  v:"min:0#分页号码错误" json:"page" description:"页数"` // 分页号码
	Size   int `d:"10" v:"max:50#分页数量最大50条" json:"size" description:"每页条数"`
}
type ElectionGetRes struct {
	*model.ElectionGetOut
}
