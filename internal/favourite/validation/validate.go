package validation

import (
	"gin-restapi/internal/model"

	"gorm.io/gorm"
)

func CheckDuplicateDish() error {
	var user model.User
	var dish model.Dish

	for _, d := range user.Dishes {
		if d.ID == dish.ID {
			return gorm.ErrDuplicatedKey
		}
	}

	return nil
}
