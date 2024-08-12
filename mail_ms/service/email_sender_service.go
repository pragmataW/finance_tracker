package service

import (
	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"github.com/pragmataW/finance_tracker/mail_ms/dto"
)

func (e *EmailSenderService) SendEmail(email dto.Email) error {
	mailjetClient := mailjet.NewMailjetClient(e.mailjetApiKey, e.mailjetSecretKey)
	messages := mailjet.MessagesV31{
		Info: []mailjet.InfoMessagesV31{
			{
				From: &mailjet.RecipientV31{
					Email: email.FromMail,
					Name:  email.FromName,
				},
				To: &mailjet.RecipientsV31{
					{
						Email: email.ToMail,
						Name:  email.ToName,
					},
				},
				Subject:  email.Subject,
				HTMLPart: email.Message,
			},
		},
	}

	_, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		e.logger.Error(err.Error())
		return err
	}

	e.logger.Info("Email sent successfully")
	return nil
}
