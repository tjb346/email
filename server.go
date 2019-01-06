package email

import (
	"net/smtp"
	"strconv"
)

type MailServer interface {
	SendMail(mail EMail) error
}

type SMTPServer struct {
	Host     string
	Port     int
	Username string
	Password string
}

func (server *SMTPServer) ServerName() string {
	return server.Host + ":" + strconv.Itoa(server.Port)
}

func (server *SMTPServer) SendMail(mail EMail) error {
	auth := smtp.PlainAuth("", server.Username, server.Password, server.Host)

	receivers := append(mail.To, mail.Cc...)
	receivers = append(receivers, mail.Bcc...)

	return smtp.SendMail(server.ServerName(), auth, mail.From, receivers, []byte(mail.BuildMessage()))
}
