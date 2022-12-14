package v1

import (
	"election/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 创建用户登录
type UserElectionGetReq struct {
	g.Meta `path:"/election" method:"get" tags:"UserService" summary:"election get"`
}
type UserElectionGetRes struct {
	*model.UserElectionGetOut
}

// 用户投票
type UserVoteReq struct {
	g.Meta      `path:"/vote" method:"post" tags:"UserService" summary:"election create"`
	IdCard      string `v:"regex:[a-zA-Z]{1}[0-9]{6}\\([0-9]{1}\\)#身份证格式错误" json:"idCard" description:"用户身份证"`
	Email       string `v:"email" json:"email" description:"用户邮箱"`
	CandidateId int    `v:"required" json:"candidateId" description:"候选人id"`
	ElectionId  int    `v:"required" json:"electionId" description:"选举id"`
}
type UserVoteRes struct{}
