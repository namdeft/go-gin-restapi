package services

import (
	"context"

	"gin-restapi/internal/favourite/repository"
	"gin-restapi/internal/model"
)

type FavouriteService interface {
	AddFavourite(ctx context.Context, userId int, dishId int) error
	GetFavourites(ctx context.Context, userId int) ([]model.Dish, error)
	DeleteFavourite(ctx context.Context, userId int, dishId int) error
}

type favouriteService struct {
	favouriteRepository repository.FavouriteRepository
}

func NewFavouriteService(favouriteRepo repository.FavouriteRepository) *favouriteService {
	return &favouriteService{favouriteRepository: favouriteRepo}
}

func (service *favouriteService) AddFavourite(ctx context.Context, userId int, dishId int) error {
	if err := service.favouriteRepository.AddFavourite(userId, dishId); err != nil {
		return err
	}

	return nil
}

func (service *favouriteService) GetFavourites(ctx context.Context, userId int) ([]model.Dish, error) {
	dishes, err := service.favouriteRepository.GetFavourites(userId)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

func (service *favouriteService) DeleteFavourite(ctx context.Context, userId int, dishId int) error {
	if err := service.favouriteRepository.DeleteFavourite(userId, dishId); err != nil {
		return err
	}
	return nil
}
