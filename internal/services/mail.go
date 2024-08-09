package services

import (
	"fmt"
	"log/slog"
	"net/smtp"

	"github.com/lwinmgmg/user-go/env"
)

var (
	MailSender *MailService
)

type MailService struct {
	senderMail string
	password   string
	Host       string
	Port       int
	Enable     bool
}

func (sender *MailService) getAuth() smtp.Auth {
	return smtp.PlainAuth("", sender.senderMail, sender.password, sender.Host)
}

func (sender *MailService) Send(message string, recipient []string) error {
	slog.Info(fmt.Sprintf("Sending email to -> %v", recipient))
	if !sender.Enable {
		slog.Warn(fmt.Sprintf("Email server is not enable : To -> %v; Message -> %v", recipient, message))
		return nil
	}
	err := smtp.SendMail(fmt.Sprintf("%v:%v", sender.Host, sender.Port), sender.getAuth(), sender.senderMail, recipient, []byte(message))
	if err != nil {
		slog.Error(fmt.Sprintf("Error on sending email %v\n", err))
		return err
	}
	return nil
}

func NewMailService(mailConf *env.EmailServer) *MailService {
	return &MailService{
		senderMail: mailConf.Email,
		password:   mailConf.Password,
		Host:       mailConf.Host,
		Port:       mailConf.Port,
		Enable:     mailConf.Enable,
	}
}
