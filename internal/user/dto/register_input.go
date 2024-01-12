package dto

import (
	"gin-restapi/internal/user/validation"

	"github.com/go-playground/validator/v10"
)

type RegisterInput struct {
	Name     string `json:"name" validate:"required,gte=1,lte=20"`
	Email    string `json:"email" validate:"required,email,gte=16,lte=30"`
	Password string `json:"password" validate:"required,validatePassword"`
}

func (input *RegisterInput) ValidateRegisterInput() error {
	validate := validator.New()
	validate.RegisterValidation("validatePassword", func(fl validator.FieldLevel) bool {
		return validation.ValidatePassword(fl.Field().String())
	})
	err := validate.Struct(input)

	return err
}
