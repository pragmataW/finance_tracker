package service

import "github.com/pragmataW/finance_tracker/mail_ms/pkg"

type EmailSenderService struct{
	mailjetApiKey string
	mailjetSecretKey string
	logger *pkg.Logger
}

func New(mailjetApiKey string, mailjetSecretKey string, logger *pkg.Logger) *EmailSenderService{
	return &EmailSenderService {
		mailjetApiKey: mailjetApiKey,
		mailjetSecretKey: mailjetSecretKey,
		logger: logger,
	}
}