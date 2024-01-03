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

type DishCreation struct {
	Id    int    `json:"-"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func (DishCreation) TableName() string {
	return "dish"
}

type DishUpdation struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	Status string `json:"status"`
}

func (DishUpdation) TableName() string {
	return "dish"
}
