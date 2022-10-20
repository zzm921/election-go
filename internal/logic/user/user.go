package user

import (
	"context"
	"election/internal/consts"
	"election/internal/dao"
	"election/internal/model"
	"election/internal/model/do"
	"election/internal/model/entity"
	"election/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}
func New() *sUser {
	return &sUser{}
}

func (s *sUser) GetElection(ctx context.Context, in model.UserElectionGetInput) (*model.UserElectionGetOut, error) {
	//查询当前正在进行的选举
	var currentElection *entity.Elections
	err := dao.Elections.Ctx(ctx).Where(g.Map{"status": 1}).WithAll().Scan(&currentElection)
	if err != nil {
		return nil, err
	}
	//当前没有开始的选举
	if currentElection == nil {
		return nil, gerror.NewCode(gcode.New(consts.DataNoExistCode, "", nil), "当前没有进行中选举")
	}
	//处理list数据
	candidates := currentElection.Candidates
	newCandidateList := []*model.CandidateListObject{}
	for _, candidate := range candidates {
		newCandiateObj := model.CandidateListObject{
			CandidateId:   candidate.CandidateId,
			VoteCount:     candidate.VoteCount,
			CandidateInfo: candidate.CandidateInfo,
		}
		newCandidateList = append(newCandidateList, &newCandiateObj)
	}
	userElectionGetOut := model.UserElectionGetOut{
		Id:           currentElection.Id,
		Title:        currentElection.Title,
		Introduction: currentElection.Introduction,
		Status:       currentElection.Status,
		Candidates:   newCandidateList,
	}

	return &userElectionGetOut, nil
}

func (s *sUser) Vote(ctx context.Context, in model.UserVoteInput) (err error) {
	//查询选举是否存在
	var currentElection *entity.Elections
	err = dao.Elections.Ctx(ctx).Where(g.Map{"status": 1, "id": in.ElectionId}).Scan(&currentElection)
	if err != nil {
		return err
	}
	//找不到，表示选举不存在或已结束
	if currentElection == nil {
		return gerror.NewCode(gcode.New(consts.DataNoExistCode, "", nil), "该选举不存在或已结束")
	}
	//判断用户是否已经投过票
	var userElectionDetail *entity.ElectionConfigCandidates
	err = dao.ElectionUserDetails.Ctx(ctx).Where(g.Map{"electionId": in.ElectionId, "idCard": in.IdCard}).Scan(&userElectionDetail)
	if err != nil {
		return err
	}
	if userElectionDetail != nil {
		return gerror.NewCode(gcode.New(consts.DataExistCode, "", nil), "该用户已经投过票")
	}

	//开始投票
	dao.ElectionUserDetails.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		//将投票记录存入表中
		_, err := dao.ElectionUserDetails.Ctx(ctx).Data(do.ElectionUserDetails{
			ElectionId:  in.ElectionId,
			CandidateId: in.CandidateId,
			Email:       in.Email,
			IdCard:      in.IdCard,
		}).Insert()
		if err != nil {
			return err
		}
		//票数记录+1
		_, err = dao.ElectionConfigCandidates.Ctx(ctx).Data(do.ElectionConfigCandidates{
			VoteCount: gdb.Raw("VoteCount+1"),
		}).Where(g.Map{"ElectionId": in.ElectionId, "CandidateId": in.CandidateId}).Update()

		if err != nil {
			return err
		}
		return err

	})
	return nil
}
