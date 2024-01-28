package services

import (
	"fmt"
	"log"
	"net/smtp"
)

var (
	MailSender *MailService
)

type MailService struct {
	senderMail string
	password   string
	Host       string
	Port       int
	IsEnable   bool
}

func (sender *MailService) getAuth() smtp.Auth {
	return smtp.PlainAuth("", sender.senderMail, sender.password, sender.Host)
}

func (sender *MailService) Send(message string, recipient []string) error {
	if !sender.IsEnable {
		log.Println("Email server is not enable")
		return nil
	}
	err := smtp.SendMail(fmt.Sprintf("%v:%v", sender.Host, sender.Port), sender.getAuth(), sender.senderMail, recipient, []byte(message))
	if err != nil {
		fmt.Printf("Error on sending email %v\n", err)
		return err
	}
	return nil
}

func NewMailService(email, password, host string, port int, isEnable bool) *MailService {
	return &MailService{
		senderMail: email,
		password:   password,
		Host:       host,
		Port:       port,
		IsEnable:   isEnable,
	}
}
