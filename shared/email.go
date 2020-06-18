package shared

import (
	"bytes"
	"net/smtp"
	"strings"
)

type EmailSender struct {
	User     string
	Password string
	Host     string
}

type Email struct {
	Subject    string
	Recipients []string
	Content    string
	IsHtml     bool
}

func (mail Email) ContentType() string {
	if mail.IsHtml {
		return "Content-Type: text/html; charset=UTF-8"
	}
	return "Content-Type: text/plain; charset=UTF-8"
}

func (sender *EmailSender) Send(mail Email) error {
	hp := strings.Split(sender.Host, ":")
	auth := smtp.PlainAuth("", sender.User, sender.Password, hp[0])

	msg := bytes.Buffer{}
	msg.WriteString("To: " + strings.Join(mail.Recipients, ";"))
	msg.WriteString("\r\n")
	msg.WriteString("From: " + sender.User)
	msg.WriteString("\r\n")
	msg.WriteString("Subject: " + mail.Subject)
	msg.WriteString("\r\n")
	msg.WriteString(mail.ContentType())
	msg.WriteString("\r\n\r\n")
	msg.WriteString(mail.Content)

	// fmt.Println(string(msg.Bytes()))

	return smtp.SendMail(sender.Host, auth, sender.User, mail.Recipients, msg.Bytes())
}
