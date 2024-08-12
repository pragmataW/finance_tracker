package controller

import (
	"github.com/pragmataW/finance_tracker/mail_ms/dto"
	"github.com/pragmataW/finance_tracker/mail_ms/pkg"
)

type IEmailSenderService interface {
	SendEmail(email dto.Email) error
}

type EmailSenderController struct {
	service IEmailSenderService
	logger *pkg.Logger
}

func New(service IEmailSenderService, logger *pkg.Logger) *EmailSenderController{
	return &EmailSenderController{
		service: service,
		logger: logger,
	}
}
