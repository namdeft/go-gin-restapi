package dto

type DishCreation struct {
	Id    int    `json:"-"`
	Name  string `json:"name" validate:"required,gte=1,lte=30"`
	Price string `json:"price" validate:"required"`
}
