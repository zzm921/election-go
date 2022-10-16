package election

import (
	"context"
	"election/internal/dao"
	response "election/internal/lib"
	"election/internal/model"
	"election/internal/model/do"
	"election/internal/model/entity"
	"election/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	sElection struct{}
)

func init() {
	service.RegisterElection(New())
}

func New() *sElection {
	return &sElection{}
}

func (s *sElection) Create(ctx context.Context, in model.ElectionCreateInput) *response.ResultRes {
	//查看传入的候选人信息是否存在
	candidates := in.Candidates

	dao.Elections.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		creatInfo, err := dao.Elections.Ctx(ctx).Data(do.Elections{
			Title:        in.Title,
			Introduction: in.Introduction,
		}).Insert()
		electionId, err := creatInfo.LastInsertId()
		var electionConfigCandidateList []entity.ElectionConfigCandidates
		for _, candidateId := range candidates {
			electionConfigCandidateList = append(electionConfigCandidateList, entity.ElectionConfigCandidates{
				ElectionId:  int(electionId),
				CandidateId: candidateId,
				VoteCount:   0,
			})
		}
		dao.ElectionConfigCandidates.Ctx(ctx).Data(electionConfigCandidateList).Insert()
		return err

	})
	dao.Elections.Ctx(ctx).Data(do.Elections{Title: in.Title}).Insert()
	//查询是否有对应的账号密码

	return &response.ResultRes{Code: response.OkCode, Message: response.OkMsg}
}
