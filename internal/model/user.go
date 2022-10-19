package model

import "election/internal/model/entity"

// 获取当前选举信息
type UserElectionGetInput struct {
}

type CandidateListObject struct {
	CandidateId   int                `json:"candidateId"          description:"候选人id"`
	VoteCount     int                `json:"voteCount"          description:"选票"`
	CandidateInfo *entity.Candidates `json:"candidateInfo"          description:"候选人信息"`
}

type UserElectionGetOut struct {
	Id           int                    `json:"id"          description:"选举id"`
	Title        string                 `json:"title"          description:"选举标题"`
	Introduction string                 `json:"introduction"          description:"选举介绍"`
	Status       int                    `json:"status"          description:"选举状态"`
	Candidates   []*CandidateListObject `json:"candidates"          description:"候选人列表"`
}

type UserVoteInput struct {
	IdCard      string
	Email       string
	CandidateId int
	ElectionId  int
}

type UserVoteOut struct {
}
