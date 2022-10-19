package controller

import (
	"context"
	"election/internal/model"
	"election/internal/service"

	v1 "election/api/v1"
)

var Candidate = cCandidate{}

type cCandidate struct{}

func (c *cCandidate) Create(ctx context.Context, req *v1.CandidateCreateReq) (res *v1.CandidateCreateRes, err error) {
	err = service.Candidate().Create(ctx, model.CandidateCreateInput{
		Name:         req.Name,
		Introduction: req.Introduction,
		Image:        req.Image,
	})
	return
}

func (c *cCandidate) Update(ctx context.Context, req *v1.CandidateUpdateReq) (res *v1.CandidateUpdateRes, err error) {
	err = service.Candidate().Update(ctx, model.CandidateUpdateInput{
		CandidateId:  req.CandidateId,
		Name:         req.Name,
		Introduction: req.Introduction,
		Image:        req.Image,
	})
	return
}

func (c *cCandidate) Get(ctx context.Context, req *v1.CandidateGetReq) (res *v1.CandidateGetRes, err error) {
	candidateGetOut, err := service.Candidate().Get(ctx, model.CandidateGetInput{
		Page: req.Page,
		Size: req.Size,
	})

	if err != nil {
		return nil, err
	}
	return &v1.CandidateGetRes{CandidateGetOut: candidateGetOut}, nil
}

func (c *cCandidate) ChangeStatus(ctx context.Context, req *v1.CandidateChangeStatusReq) (res *v1.CandidateChangeStatusRes, err error) {
	err = service.Candidate().ChangeStatus(ctx, model.CandidateChangeStatuInput{
		CandidateId: req.CandidateId,
		Status:      req.Status,
	})
	return
}
