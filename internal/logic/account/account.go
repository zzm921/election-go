package user

import (
	"context"
	"crypto/md5"
	"election/internal/dao"
	response "election/internal/lib"
	"election/internal/model"
	"election/internal/model/do"
	"election/internal/model/entity"
	"election/internal/service"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/frame/g"
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
func (s *sAccount) Login(ctx context.Context, in model.AccountLoginInput) *response.ResultRes {
	var account *entity.Accounts
	//查询是否有对应的账号密码
	var err = dao.Accounts.Ctx(ctx).Where(do.Accounts{
		Username: in.Username,
		Password: in.Password,
	}).Scan(&account)
	if err != nil {
		return &response.ResultRes{Code: response.ErrorCode, Message: response.ErrorMsg, Data: err}
	}
	if account == nil {
		return &response.ResultRes{Code: response.AccountValidErrCode, Message: response.AccountValidErrMsg}
	}
	//登录成功，生成token，将数据存入redis
	timeNow := time.Now().UnixMilli()
	data := []byte(strconv.FormatInt(timeNow, 10) + account.Username)
	has := md5.Sum(data)
	//根据时间戳，账号名称生成token
	token := fmt.Sprintf("%x", has) //将[]byte转成16进制
	accountJsonStr, _ := json.Marshal(&account)
	//将数据存入redis
	_, err = g.Redis().Do(ctx, "SETEX", "token", 86400, string(accountJsonStr))
	if err != nil {
		return &response.ResultRes{Code: response.ErrorCode, Message: response.ErrorMsg, Data: err}
	}
	loginOut := model.AccountLoginOut{
		Token:    token,
		Username: account.Username,
		Role:     account.Role,
	}
	return &response.ResultRes{Code: response.OkCode, Message: response.OkMsg, Data: loginOut}
}
