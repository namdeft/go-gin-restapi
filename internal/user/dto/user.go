package dto

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type RegisterInput struct {
	Name     string `json:"name" validate:"required,gte=1,lte=20"`
	Email    string `json:"email" validate:"required,email,gte=16,lte=30"`
	Password string `json:"password" validate:"required,validatePassword"`
}

func (RegisterInput) TableName() string {
	return "user"
}

type LoginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (LoginInput) TableName() string {
	return "user"
}

func validatePassword(password string) bool {
	return len(password) >= 7 &&
		contains(password, "[0-9]") &&
		contains(password, "[A-Z]") &&
		contains(password, `[!@#$%^&*()_+{}|:"<>?~]`)
}

func contains(s, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}

func (input *RegisterInput) ValidateRegisterInput() error {
	validate := validator.New()
	validate.RegisterValidation("validatePassword", func(fl validator.FieldLevel) bool {
		return validatePassword(fl.Field().String())
	})
	err := validate.Struct(input)

	return err
}

func (input *LoginInput) ValidateLoginInput() error {
	validate := validator.New()
	err := validate.Struct(input)

	return err
}
