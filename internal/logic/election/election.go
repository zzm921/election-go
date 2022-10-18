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
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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

func (s *sElection) Create(ctx context.Context, in model.ElectionCreateInput) (err error) {

	candidates := in.Candidates
	if len(candidates) < 2 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "候选人数应该大于2")
	}
	//查看传入的候选人信息是否存在
	var candidateDbinfos = []*entity.Candidates{}
	candidateErr := dao.Candidates.Ctx(ctx).WhereIn("id", candidates).Scan(&candidateDbinfos)
	if candidateErr != nil {
		return err
	}
	if len(candidates) != len(candidateDbinfos) {
		return gerror.NewCode(gcode.CodeInvalidParameter, "非法的候选人信息")
	}
	dao.Elections.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		//创建选举并获取创建后的id
		electionId, err := dao.Elections.Ctx(ctx).Data(do.Elections{
			Title:        in.Title,
			Introduction: in.Introduction,
		}).InsertAndGetId()
		if err != nil {
			return err
		}
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
	//查询是否有对应的账号密码
	return
}

func (s *sElection) ChangeStatus(ctx context.Context, in model.ElectionChangeStatuInput) (err error) {
	electionId := in.ElectionId
	status := in.Status
	//查看该选举是否存在
	var electionsDbResult *entity.Elections
	dberr := dao.Elections.Ctx(ctx).Where(do.Elections{Id: electionId}).Scan(&electionsDbResult)
	if dberr != nil {
		return dberr
	}
	//选举不存在返回错误
	if electionsDbResult == nil {
		return gerror.NewCode(gcode.New(response.DataNoExistCode, "", nil), "选举不存在")
	}
	//校验选举状态 只能从 未开始0->开始1->结束2
	if status-electionsDbResult.Status != 1 {
		return gerror.Newf(`无法设置该status参数`)
	}
	//更新选举状态
	_, err2 := dao.Elections.Ctx(ctx).Data(g.Map{"status": status}).Where(g.Map{"Id": electionId}).Update()
	if err2 != nil {
		return err2
	}
	return
}

func (s *sElection) Get(ctx context.Context, in model.ElectionGetInput) (*model.ElectionGetOut, error) {
	page := in.Page
	limit := in.Size
	offset := (page - 1) * limit
	electionsDbResult := []*entity.Elections{}
	err := dao.Elections.Ctx(ctx).WithAll().Limit(offset, limit).Scan(&electionsDbResult)
	count, err2 := dao.Elections.Ctx(ctx).Count()
	if err != nil || err2 != nil {
		return nil, err
	}

	//处理list数据
	list := []*model.ElectionGetOutListObject{}
	for _, election := range electionsDbResult {
		candidates := election.Candidates
		newCandidate := []int{}
		for _, candidate := range candidates {
			newCandidate = append(newCandidate, candidate.CandidateId)
		}
		obj := model.ElectionGetOutListObject{
			Id:           election.Id,
			Title:        election.Title,
			Introduction: election.Introduction,
			Status:       election.Status,
			Candidates:   newCandidate,
		}
		list = append(list, &obj)
	}
	electionGetOut := model.ElectionGetOut{
		Count: count,
		List:  list,
	}

	return &electionGetOut, nil
}
