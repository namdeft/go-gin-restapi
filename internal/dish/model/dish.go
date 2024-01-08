package model

import (
	"time"
)

type Dish struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Price      string     `json:"price"`
	Status     DishStatus `json:"status"`
	Updated_At time.Time  `json:"updated_at"`
	Created_At time.Time  `json:"created_at"`
	Deleted_At time.Time  `json:"deleted_at"`
}

func (Dish) TableName() string {
	return "dish"
}
