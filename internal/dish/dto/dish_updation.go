package dto

type DishUpdation struct {
	Name   string `json:"name" validate:"required,gte=1,lte=30"`
	Price  string `json:"price" validate:"required"`
	Status string `json:"status" validate:"required,eq=unavailable|eq=available|eq=deleted"`
}
