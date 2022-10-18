package lib

import (
	"crypto/tls"
	"net/smtp"

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
func SendEmail(emailBody *EmailBody) error {
	e := email.NewEmail()
	e.From = From
	e.To = emailBody.To
	e.Subject = emailBody.Subject
	e.HTML = []byte(emailBody.HTML)
	err := e.SendWithTLS(Addr, smtp.PlainAuth("", User, Pass, Host),
		&tls.Config{InsecureSkipVerify: true, ServerName: Host})
	if err != nil {
		return err
	}
	return nil
}
