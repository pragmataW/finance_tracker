package dto

type Email struct {
	FromMail string `json:"from_mail" validate:"required,email"`
	FromName string `json:"from_name" validate:"required"`
	ToMail   string `json:"to_mail" validate:"required,email"`
	ToName   string `json:"to_name" validate:"required"`
	Message  string `json:"message" validate:"required"`
	Subject  string `json:"subject" validate:"required"`
}
