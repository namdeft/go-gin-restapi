package repository

import (
	"context"

	"gin-restapi/internal/favourite/validation"
	"gin-restapi/internal/model"

	"gorm.io/gorm"
)

type FavouriteRepository interface {
	AddFavourite(ctx context.Context, userId int, dishId int) error
	GetFavourites(ctx context.Context, userId int) ([]model.Dish, error)
	DeleteFavourite(ctx context.Context, userId int, dishId int) error
}

type favouriteConnection struct {
	db *gorm.DB
}

func NewFavouriteRepository(db *gorm.DB) FavouriteRepository {
	return &favouriteConnection{
		db: db,
	}
}

func (s *favouriteConnection) AddFavourite(ctx context.Context, userId int, dishId int) error {
	var user model.User
	if err := s.db.First(&user, userId).Error; err != nil {
		return err
	}

	var dish model.Dish
	if err := s.db.Preload("Users").First(&dish, dishId).Error; err != nil {
		return err
	}

	if err := validation.CheckDuplicateDish(); err != nil {
		return err
	}

	user.Dishes = append(user.Dishes, dish)

	if err := s.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (s *favouriteConnection) GetFavourites(ctx context.Context, userId int) ([]model.Dish, error) {
	var user model.User
	if err := s.db.Preload("Dishes").First(&user, userId).Error; err != nil {
		return nil, err
	}

	return user.Dishes, nil
}

func (s *favouriteConnection) DeleteFavourite(ctx context.Context, userId int, dishId int) error {
	var user model.User
	if err := s.db.First(&user, userId).Error; err != nil {
		return err
	}

	if err := s.db.Model(&user).Association("Dishes").Delete(&model.Dish{ID: dishId}); err != nil {
		return err
	}

	return nil
}
