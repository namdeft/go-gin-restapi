package dto

import "github.com/go-playground/validator/v10"

type DishCreation struct {
	Id    int    `json:"-"`
	Name  string `json:"name" validate:"required,gte=1,lte=30"`
	Price string `json:"price" validate:"required"`
}

func (DishCreation) TableName() string {
	return "dish"
}

type DishUpdation struct {
	Name   string `json:"name" validate:"required,gte=1,lte=30"`
	Price  string `json:"price" validate:"required"`
	Status string `json:"status" validate:"required,eq=unavailable|eq=available|eq=deleted"`
}

func (DishUpdation) TableName() string {
	return "dish"
}

func (input *DishCreation) ValidateDishCreation() error {
	validate := validator.New()
	err := validate.Struct(input)

	return err
}

func (input *DishUpdation) ValidateDishUpdation() error {
	validate := validator.New()
	err := validate.Struct(input)

	return err
}
