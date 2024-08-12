package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pragmataW/finance_tracker/mail_ms/dto"
)

func (e *EmailSenderController) SendEmail(c *fiber.Ctx) error {
	var body dto.Email
	if err := c.BodyParser(&body); err != nil{
		e.logger.Error("error while parsing body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil{
		e.logger.Error("error while validating body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := e.service.SendEmail(body); err != nil{
		e.logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "email not sent",
		})
	}
	
	e.logger.Info("email sent")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "email sent",
	})
}