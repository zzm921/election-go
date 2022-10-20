package controller

import (
	"context"
	"election/internal/model"
	"election/internal/service"

	v1 "election/api/v1"
)

var Electoin = cElection{}

type cElection struct{}

func (c *cElection) Create(ctx context.Context, req *v1.ElectionCreateReq) (res *v1.ElectionCreateRes, err error) {
	err = service.Election().Create(ctx, model.ElectionCreateInput{
		Title:        req.Title,
		Introduction: req.Introduction,
		Candidates:   req.Candidates,
	})
	return
}

func (c *cElection) Get(ctx context.Context, req *v1.ElectionGetReq) (res *v1.ElectionGetRes, err error) {
	electionGetOut, err := service.Election().Get(ctx, model.ElectionGetInput{
		Page: req.Page,
		Size: req.Size,
	})

	if err != nil {
		return nil, err
	}
	return &v1.ElectionGetRes{ElectionGetOut: electionGetOut}, nil
}

func (c *cElection) ChangeStatus(ctx context.Context, req *v1.ElectionChangeStatusReq) (res *v1.ElectionChangeStatusRes, err error) {
	err = service.Election().ChangeStatus(ctx, model.ElectionChangeStatuInput{
		ElectionId: req.ElectionId,
		Status:     req.Status,
	})
	return
}

func (c *cElection) GetElectionCandidate(ctx context.Context, req *v1.ElectionCandidateGetReq) (res *v1.ElectionCandidateGetRes, err error) {
	electionCandidateGetOut, err := service.Election().GetElectionCandidate(ctx, model.ElectionCandidateGetInput{
		Page:       req.Page,
		Size:       req.Size,
		ElectionId: req.ElectionId,
	})

	if err != nil {
		return nil, err
	}
	return &v1.ElectionCandidateGetRes{ElectionCandidateGetOut: electionCandidateGetOut}, nil
}

func (c *cElection) GetElectionCandidateVote(ctx context.Context, req *v1.ElectionCandidateVoteGetReq) (res *v1.ElectionCandidateVoteGetRes, err error) {
	electionCandidateVoteGetOut, err := service.Election().GetElectionCandidateVote(ctx, model.ElectionCandidateVoteGetInput{
		Page:        req.Page,
		Size:        req.Size,
		ElectionId:  req.ElectionId,
		CandidateId: req.CandidateId,
	})

	if err != nil {
		return nil, err
	}
	return &v1.ElectionCandidateVoteGetRes{ElectionCandidateVoteGetOut: electionCandidateVoteGetOut}, nil
}
