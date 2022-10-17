package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)
//创建用户登录
type UserLoginReq struct {
	g.Meta       `path:"/user/login" method:"post" tags:"UserService" summary:"election create"`
	IdCard        string `v:"required"`
}
type UserLoginRes struct{}



//用户投票
type UserVoteReq struct {
	g.Meta       `path:"/user/vote" method:"post" tags:"UserService" summary:"election create"`
	IdCard        string `v:"required"`
	Email        string `v:"required"`
	CandidateId        int `v:"required"`
}
type UserVoteRes struct{}
