package services

import (
	"context"
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/dto"
	"gin-restapi/internal/dish/model"
	"gin-restapi/internal/dish/repository"
)

type DishService interface {
	CreateNewDish(ctx context.Context, data *dto.DishCreation) error
	GetDishes(ctx context.Context, paging *common.Paging) ([]model.Dish, error)
	GetDish(ctx context.Context, id int) (*model.Dish, error)
	UpdateDish(ctx context.Context, id int, data *dto.DishUpdation) error
	DeleteDish(ctx context.Context, id int) error
}

type dishService struct {
	dishRepository repository.DishRepository
}

func NewDishService(dishRepo repository.DishRepository) *dishService {
	return &dishService{dishRepository: dishRepo}
}

func (service *dishService) CreateNewDish(ctx context.Context, data *dto.DishCreation) error {
	if err := service.dishRepository.CreateDish(ctx, data); err != nil {
		return err
	}

	return nil
}

func (service *dishService) GetDishes(ctx context.Context, paging *common.Paging) ([]model.Dish, error) {
	dishes, err := service.dishRepository.GetDishes(ctx, paging)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

func (service *dishService) GetDish(ctx context.Context, id int) (*model.Dish, error) {
	dish, err := service.dishRepository.GetDish(ctx, id)
	if err != nil {
		return nil, err
	}

	return dish, nil
}

func (service *dishService) UpdateDish(ctx context.Context, id int, data *dto.DishUpdation) error {
	if err := service.dishRepository.UpdateDish(ctx, id, data); err != nil {
		return err
	}

	return nil
}

func (service *dishService) DeleteDish(ctx context.Context, id int) error {
	if err := service.dishRepository.DeleteDish(ctx, id); err != nil {
		return err
	}

	return nil
}
