package service

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pragmataW/finance_tracker/mail_ms/dto"
	"github.com/stretchr/testify/assert"
)

func setUpTest() error {
	return godotenv.Load("../.env")
}

func TestSendEmail(t *testing.T) {
	setUpTest()

	apikey := os.Getenv("MAILJET_API_KEY")
	secretKey := os.Getenv("MAILJET_SECRET_KEY")
	assert.NotEmpty(t, apikey)
	assert.NotEmpty(t, secretKey)

	emailSender := EmailSenderService{
		mailjetApiKey:    apikey,
		mailjetSecretKey: secretKey,
	}

	email := dto.Email{
		FromMail: "ciftciyusuf700@gmail.com",
		FromName: "PragmaTracker",
		ToMail:   "ciftciyusuf700@gmail.com",
		ToName:   "yusuf",
		Subject:  "Test Subject",
		Message:  "Test Message",
	}

	err := emailSender.SendEmail(email)
	assert.NoError(t, err)
}

func TestSendEmailButWithError(t *testing.T) {
	setUpTest()

	apikey := os.Getenv("MAILJET_API_KEY")
	secretKey := os.Getenv("MAILJET_SECRET_KEY")
	assert.NotEmpty(t, apikey)
	assert.NotEmpty(t, secretKey)

	emailSender := EmailSenderService{
		mailjetApiKey:    apikey,
		mailjetSecretKey: secretKey,
	}

	email := dto.Email{
		FromName: "PragmaTracker",
		ToMail:   "ciftciyusuf700@gmail.com",
		ToName:   "yusuf",
		Subject:  "Test Subject",
		Message:  "Test Message",
	}

	err := emailSender.SendEmail(email)
	assert.Error(t, err)
}
