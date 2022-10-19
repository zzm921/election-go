package controller

import (
	"context"
	"election/internal/model"
	"election/internal/service"

	v1 "election/api/v1"
)

var User = cUser{}

type cUser struct{}

func (c *cUser) GetElection(ctx context.Context, req *v1.UserElectionGetReq) (res *v1.UserElectionGetRes, err error) {
	userElectionGetOut, err := service.User().GetElection(ctx, model.UserElectionGetInput{})
	if err != nil {
		return nil, err
	}
	return &v1.UserElectionGetRes{UserElectionGetOut: userElectionGetOut}, nil
}

func (c *cUser) Vote(ctx context.Context, req *v1.UserVoteReq) (res *v1.UserVoteRes, err error) {
	err = service.User().Vote(ctx, model.UserVoteInput{
		IdCard:      req.IdCard,
		Email:       req.Email,
		CandidateId: req.CandidateId,
		ElectionId:  req.ElectionId,
	})
	return
}
