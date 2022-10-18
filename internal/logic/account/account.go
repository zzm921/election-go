package user

import (
	"context"
	"crypto/md5"
	"election/internal/dao"
	"election/internal/model"
	"election/internal/model/do"
	"election/internal/model/entity"
	"election/internal/service"
	"fmt"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	sAccount struct{}
)

func init() {
	service.RegisterAccount(New())
}

func New() *sAccount {
	return &sAccount{}
}

// SignIn creates session for given user account.
func (s *sAccount) Login(ctx context.Context, in model.AccountLoginInput) (*model.AccountLoginOut, error) {
	var account *entity.Accounts
	//查询是否有对应的账号密码
	var err = dao.Accounts.Ctx(ctx).Where(do.Accounts{
		Username: in.Username,
		Password: in.Password,
	}).Scan(&account)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "账号或密码错误")
	}
	//登录成功，生成token，将数据存入redis
	timeNow := time.Now().UnixMilli()
	data := []byte(strconv.FormatInt(timeNow, 10) + account.Username)
	has := md5.Sum(data)
	//根据时间戳，账号名称生成token
	token := fmt.Sprintf("%x", has) //将[]byte转成16进制

	loginOut := model.AccountLoginOut{
		Token:    token,
		Id:       account.Id,
		Username: account.Username,
		Role:     account.Role,
	}
	return &loginOut, nil
}
