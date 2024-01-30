package repository

import (
	"gin-restapi/internal/model"

	"gorm.io/gorm"
)

type FavouriteRepository interface {
	AddFavourite(userId int, dishId int) error
	GetFavourites(userId int) ([]model.Dish, error)
	DeleteFavourite(userId int, dishId int) error
	CheckDishExists(dishId int) error
}

type favouriteConnection struct {
	db *gorm.DB
}

func NewFavouriteRepository(db *gorm.DB) FavouriteRepository {
	return &favouriteConnection{
		db: db,
	}
}

func (s *favouriteConnection) AddFavourite(userId int, dishId int) error {
	var user model.User
	if err := s.db.Preload("Dishes").First(&user, userId).Error; err != nil {
		return err
	}

	var dish model.Dish

	user.Dishes = append(user.Dishes, dish)

	if err := s.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (s *favouriteConnection) GetFavourites(userId int) ([]model.Dish, error) {
	var user model.User
	if err := s.db.Preload("Dishes").First(&user, userId).Error; err != nil {
		return nil, err
	}

	return user.Dishes, nil
}

func (s *favouriteConnection) DeleteFavourite(userId int, dishId int) error {
	var user model.User
	if err := s.db.First(&user, userId).Error; err != nil {
		return err
	}

	if err := s.db.Model(&user).Association("Dishes").Delete(&model.Dish{ID: dishId}); err != nil {
		return err
	}

	return nil
}

func (s *favouriteConnection) CheckDishExists(dishId int) error {
	var dish model.Dish
	if err := s.db.First(&dish, dishId).Error; err != nil {
		return err
	}

	return nil
}
