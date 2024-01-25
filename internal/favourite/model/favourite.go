package model

import (
	"gin-restapi/internal/dish/model"
	"time"
)

type Dish struct {
	ID         int              `json:"id"`
	Name       string           `json:"name"`
	Price      string           `json:"price"`
	Status     model.DishStatus `json:"status"`
	Updated_At time.Time        `json:"updated_at"`
	Created_At time.Time        `json:"created_at"`
	Deleted_At time.Time        `json:"deleted_at"`
	Users      []User           `gorm:"many2many:favourite" json:"-"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Dishes   []Dish `gorm:"many2many:favourite" json:"-"`
}
