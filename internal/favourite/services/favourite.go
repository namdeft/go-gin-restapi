package services

import (
	"context"
	"gin-restapi/internal/favourite/model"
	"gin-restapi/internal/favourite/repository"
)

type FavouriteService interface {
	AddFavourite(ctx context.Context, userId int, dishId int) error
	GetFavourites(ctx context.Context, userId int) ([]model.Dish, error)
	DeleteFavourite(ctx context.Context, userId int, dishId int) error
}

type favouriteService struct {
	favouriteRepository repository.FavouriteRepository
}

func NewDishService(favouriteRepo repository.FavouriteRepository) *favouriteService {
	return &favouriteService{favouriteRepository: favouriteRepo}
}

func (service *favouriteService) AddFavourite(ctx context.Context, userId int, dishId int) error {
	if err := service.favouriteRepository.AddFavourite(ctx, userId, dishId); err != nil {
		return err
	}

	return nil
}

func (service *favouriteService) GetFavourites(ctx context.Context, userId int) ([]model.Dish, error) {
	dishes, err := service.favouriteRepository.GetFavourites(ctx, userId)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

func (service *favouriteService) DeleteFavourite(ctx context.Context, userId int, dishId int) error {
	if err := service.favouriteRepository.DeleteFavourite(ctx, userId, dishId); err != nil {
		return err
	}
	return nil
}
