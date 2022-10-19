package model

import (
	"election/internal/model/entity"
)

// 创建选举
type ElectionCreateInput struct {
	Title        string
	Introduction string
	Candidates   []int
}

type ElectionCreateOut struct {
}

// 修改选举状态
type ElectionChangeStatuInput struct {
	ElectionId int
	Status     int
}

type ElectionChangeStatuOut struct {
}

// 修改选举状态
type ElectionGetInput struct {
	Page int
	Size int
}

type ElectionGetOutListObject struct {
	Id           int    `json:"id"          description:"选举id"`
	Title        string `json:"title"          description:"选举标题"`
	Introduction string `json:"introduction"          description:"选举介绍"`
	Status       int    `json:"status"          description:"选举状态 0-未开始 1-进行中 2-已结束"`
	Candidates   []int  `json:"candiates"          description:"候选人id列表"`
}

type ElectionGetOut struct {
	Count int                         `json:"count"          description:"总数"`
	List  []*ElectionGetOutListObject `json:"list"          description:"数据"`
}

// 获取选举所有候选人信息
type ElectionCandidateGetInput struct {
	Page       int
	Size       int
	ElectionId int
}

type ElectionCandidateGetOutListObject struct {
	CandidateId   int                `json:"candidateId"          description:"候选人id"`
	VoteCount     int                `json:"voteCount"          description:"候选人当选票数"`
	CandidateInfo *entity.Candidates `json:"candidateInfo"          description:"候选人信息"`
}

type ElectionCandidateGetOut struct {
	Count int                                  `json:"count"          description:"总数"`
	List  []*ElectionCandidateGetOutListObject `json:"list"          description:"数据"`
}

// 获取选举所有候选人信息
type ElectionCandidateVoteGetInput struct {
	Page        int
	Size        int
	ElectionId  int
	CandidateId int
}

type ElectionCandidateVoteGetOut struct {
	Count int                           `json:"count"          description:"总数"`
	List  []*entity.ElectionUserDetails `json:"list"          description:"数据"`
}
