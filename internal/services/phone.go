package services

import (
	"fmt"
	"log/slog"
)

type PhoneService struct {
}

func (ps *PhoneService) Send(mesg string, dest ...string) error {
	slog.Info(fmt.Sprintf("Passcode %v has been sent to %v", mesg, dest))
	return nil
}

func NewPhoneServer() *PhoneService {
	return &PhoneService{}
}
