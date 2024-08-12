package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pragmataW/finance_tracker/mail_ms/controller"
	"github.com/pragmataW/finance_tracker/mail_ms/pkg"
	"github.com/pragmataW/finance_tracker/mail_ms/service"
)

func init(){
	if err := godotenv.Load("../.env"); err != nil{
		log.Fatal(err)
	}
}

func main() {
	apiKey := os.Getenv("MAILJET_API_KEY")
	secretKey := os.Getenv("MAILJET_SECRET_KEY")
	logger := pkg.NewLogger(pkg.DEBUG, os.Stdout)

	src := service.New(apiKey, secretKey, logger)
	ctrl := controller.New(src, logger)

	app := fiber.New()
	app.Post("mail/send", ctrl.SendEmail)

	if err := app.Listen(":1010"); err != nil{
		log.Fatal(err)
	}
}
