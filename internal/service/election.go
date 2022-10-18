// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"election/internal/model"
)

type (
	IElection interface {
		Create(ctx context.Context, in model.ElectionCreateInput) (err error)
		ChangeStatus(ctx context.Context, in model.ElectionChangeStatuInput) (err error)
		Get(ctx context.Context, in model.ElectionGetInput) (*model.ElectionGetOut, error)
	}
)

var (
	localElection IElection
)

func Election() IElection {
	if localElection == nil {
		panic("implement not found for interface IElection, forgot register?")
	}
	return localElection
}

func RegisterElection(i IElection) {
	localElection = i
}
