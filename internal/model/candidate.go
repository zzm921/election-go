package model

import "election/internal/model/entity"

//添加候选人
type CandidateCreateInput struct {
	Name         string
	Introduction string
	Image        string
}

type CandidateCreateOut struct {
}

type CandidateUpdateInput struct {
	CandidateId  int
	Name         string
	Introduction string
	Image        string
}

type CandidateUpdateOut struct {
}

// 修改候选人状态
type CandidateChangeStatuInput struct {
	CandidateId int
	Status      int
}

type CandidateChangeStatuOut struct {
}

// 获取
type CandidateGetInput struct {
	Page int
	Size int
}

type CandidateGetOut struct {
	Count int                  `json:"count"          description:"总数"`
	List  []*entity.Candidates `json:"list"          description:"数据"`
}
