package lib

import (
	"context"
	"crypto/tls"
	"election/internal/dao"
	"election/internal/model/entity"
	"net/smtp"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/jordan-wright/email"
)

var (
	Addr = "smtp.exmail.qq.com:465"
	Host = "smtp.exmail.qq.com"
	From = "zhangzheming@cocheer.net"
	User = "zhangzheming@cocheer.net"
	Pass = "824781943Asa"
)

type EmailBody struct {
	To      []string
	Subject string
	HTML    string
	Text    string
}

// 发送邮件
func SendEmail(emailBody interface{}) error {
	body := emailBody.(EmailBody)
	e := email.NewEmail()
	e.From = From
	e.To = body.To
	e.Subject = body.Subject
	e.HTML = []byte(body.HTML)
	err := e.SendWithTLS(Addr, smtp.PlainAuth("", User, Pass, Host),
		&tls.Config{InsecureSkipVerify: true, ServerName: Host})
	if err != nil {
		return err
	}
	return nil
}

func SendElectionResult(ctx context.Context, electionId int) {
	//定义5个协程
	//开个子协程去不停的获取邮箱
	candidates := []*entity.ElectionConfigCandidates{}
	dao.ElectionConfigCandidates.Ctx(ctx).Where(g.Map{"electionId": electionId}).WithAll().OrderDesc("voteCount").Scan(&candidates)
	html_text := "<table border=1>	<tr>	<th>候选人</th>	<th>票数</th> </tr>"
	for _, candidate := range candidates {
		html_text = html_text + "<tr>	<th>" + candidate.CandidateInfo.Name + "</th>	<th>" + strconv.Itoa(candidate.VoteCount) + "</th> </tr>"
	}
	html_text = "</table>"

	emailBody := EmailBody{
		Subject: "选举结果通知",
		HTML:    html_text,
		Text:    "选举结果通知",
	}
	page := 1
	limit := 10000
	for {
		//获取所有的投票信息
		electionUserDetails := []*entity.ElectionUserDetails{}
		offset := (page - 1) * limit
		dao.ElectionUserDetails.Ctx(ctx).Where(g.Map{"electionId": electionId}).Limit(offset, limit).Scan(&electionUserDetails)
		for _, info := range electionUserDetails {
			emailBody.To = []string{info.Email}
			var sendEmailTask = PoolTask{
				Args: emailBody,
				Do:   SendEmail,
			}
			//每个邮箱任务塞入任务通道
			PoolJobs <- sendEmailTask
		}
		if len(electionUserDetails) < limit {
			break
		}
		page++
	}
	PoolRun(2)
}
