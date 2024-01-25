package repository

import (
	"context"
	"gin-restapi/internal/favourite/model"

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
	if err := s.db.First(&dish, dishId).Error; err != nil {
		return err
	}

	for _, d := range user.Dishes {
		if d.ID == dish.ID {
			return gorm.ErrDuplicatedKey
		}
	}

	user.Dishes = append(user.Dishes, dish)

	s.db.Save(&user)

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
	if err := s.db.Preload("Dishes").First(&user, userId).Error; err != nil {
		return err
	}

	for _, dish := range user.Dishes {
		if dish.ID == dishId {
			s.db.Model(&user).Association("Dishes").Delete(&dish)
		}
	}

	return nil
}
