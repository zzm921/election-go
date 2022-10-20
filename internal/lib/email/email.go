package email

import (
	"context"
	"crypto/tls"
	"election/internal/dao"
	"election/internal/lib/pool"
	"election/internal/model/entity"
	"net/smtp"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/jordan-wright/email"
)

var (
	EmailAddr = "smtp.exmail.qq.com:465"
	EmailHost = "smtp.exmail.qq.com"
	EmailFrom = "zhangzheming@cocheer.net"
	EmailUser = "zhangzheming@cocheer.net"
	EmailPass = "824781943Qwq"
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
	e.From = EmailFrom
	e.To = body.To
	e.Subject = body.Subject
	e.HTML = []byte(body.HTML)
	err := e.SendWithTLS(EmailAddr, smtp.PlainAuth("", EmailUser, EmailPass, EmailHost),
		&tls.Config{InsecureSkipVerify: true, ServerName: EmailHost})
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
	html_text = html_text + "</table>"

	emailBody := EmailBody{
		Subject: "选举结果通知",
		HTML:    html_text,
		Text:    "选举结果通知",
	}
	page := 1
	limit := 10000
	//创建一个协程池,最大开启5个协程worker
	p := pool.NewPool(5)
	for {
		//获取所有的投票信息
		electionUserDetails := []*entity.ElectionUserDetails{}
		offset := (page - 1) * limit
		dao.ElectionUserDetails.Ctx(ctx).Where(g.Map{"electionId": electionId}).Limit(offset, limit).Scan(&electionUserDetails)
		for _, info := range electionUserDetails {
			emailBody.To = []string{info.Email}
			t := pool.NewTask(SendEmail, emailBody)
			//不断的向 Pool 输入发送邮件数据
			go func() {
				p.EntryChannel <- t
			}()
		}
		if len(electionUserDetails) < limit {
			break
		}
		page++
	}
	//启动协程池p
	p.Run()
}
