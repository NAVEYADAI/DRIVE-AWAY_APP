package LogInApiController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type sendMail struct {
}

func SendMail(c *gin.Context) {
	sender := NewEmailSender("Nave Yehiel Yadai", "naveyehielyadai@gmail.com", "socq exxn oanr libq")

	subject := "Drive Away"
	content := `
	<h1>Drive Away</h1>
	<p>This is a test message from </p>
`
	to := []string{"naveyadai@gmail.com"}
	var attachFile = []string{}
	err := sender.SendEmail(subject, content, to, nil, nil, attachFile)
	if err != nil {

	}
}

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(
		Subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
}
type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewEmailSender(name string, fromEmailAddress string, fromEmailPassword string) *GmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *GmailSender) SendEmail(
	Subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = Subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("filed to attach file %s: to %w", f, err)
		}
	}
	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)
}
