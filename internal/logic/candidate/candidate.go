package election

import (
	"context"
	"election/internal/dao"
	response "election/internal/lib"
	"election/internal/model"
	"election/internal/model/do"
	"election/internal/model/entity"
	"election/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	sCandidate struct{}
)

func init() {
	service.RegisterCandidate(New())
}

func New() *sCandidate {
	return &sCandidate{}
}

func (s *sCandidate) Create(ctx context.Context, in model.CandidateCreateInput) (err error) {
	//新增候选人
	_, err = dao.Candidates.Ctx(ctx).Data(do.Candidates{
		Name:         in.Name,
		Introduction: in.Introduction,
		Image:        in.Image,
	}).Insert()
	return err
}

func (s *sCandidate) Update(ctx context.Context, in model.CandidateUpdateInput) (err error) {
	candidateId := in.CandidateId
	//查看修改的候选人信息是否存在
	var candidateDbResult *entity.Candidates
	dberr := dao.Candidates.Ctx(ctx).Where(do.Candidates{Id: candidateId}).Scan(&candidateDbResult)
	if dberr != nil {
		return dberr
	}
	//不存在返回错误
	if candidateDbResult == nil {
		return gerror.NewCode(gcode.New(response.DataNoExistCode, "", nil), "需更新数据不存在")
	}
	//更新候选人信息
	_, err = dao.Candidates.Ctx(ctx).Data(do.Candidates{
		Name:         in.Name,
		Introduction: in.Introduction,
		Image:        in.Image,
	}).Where(g.Map{"Id": candidateId}).Update()

	return err
}

func (s *sCandidate) ChangeStatus(ctx context.Context, in model.CandidateChangeStatuInput) (err error) {
	candidateId := in.CandidateId
	status := in.Status
	//查看修改的候选人信息是否存在
	var candidateDbResult *entity.Candidates
	dberr := dao.Candidates.Ctx(ctx).Where(do.Candidates{Id: candidateId}).Scan(&candidateDbResult)
	if dberr != nil {
		return dberr
	}
	//不存在返回错误
	if candidateDbResult == nil {
		return gerror.NewCode(gcode.New(response.DataNoExistCode, "", nil), "需更新数据不存在")
	}
	//更新候选人状态
	_, err = dao.Candidates.Ctx(ctx).Data(g.Map{"status": status}).Where(g.Map{"Id": candidateId}).Update()
	return err
}

func (s *sCandidate) Get(ctx context.Context, in model.CandidateGetInput) (*model.CandidateGetOut, error) {
	page := in.Page
	limit := in.Size
	offset := (page - 1) * limit
	candidatesDbResult := []*entity.Candidates{}
	//分页查询
	err := dao.Candidates.Ctx(ctx).Limit(offset, limit).Scan(&candidatesDbResult)
	//查询总数
	count, countErr := dao.Candidates.Ctx(ctx).Count()
	if err != nil {
		return nil, err
	}
	if countErr != nil {
		return nil, countErr
	}
	candidatesGetOut := model.CandidateGetOut{
		Count: count,
		List:  candidatesDbResult,
	}

	return &candidatesGetOut, nil
}
