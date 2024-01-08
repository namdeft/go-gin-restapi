package services

import (
	"context"
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/dto"
	"gin-restapi/internal/dish/model"
)

type DishStorage interface {
	GetDishes(ctx context.Context, paging *common.Paging) ([]model.Dish, error)
	GetDish(ctx context.Context, id int) (*model.Dish, error)
	CreateDish(ctx context.Context, data *dto.DishCreation) error
	UpdateDish(ctx context.Context, id int, data *dto.DishUpdation) error
	DeleteDish(ctx context.Context, id int) error
}

type dishService struct {
	store DishStorage
}

func DishService(store DishStorage) *dishService {
	return &dishService{store: store}
}

func (business *dishService) CreateNewDish(ctx context.Context, data *dto.DishCreation) error {
	if err := business.store.CreateDish(ctx, data); err != nil {
		return err
	}

	return nil
}

func (service *dishService) GetDishes(ctx context.Context, paging *common.Paging) ([]model.Dish, error) {
	dishes, err := service.store.GetDishes(ctx, paging)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

func (service *dishService) GetDish(ctx context.Context, id int) (*model.Dish, error) {
	dish, err := service.store.GetDish(ctx, id)
	if err != nil {
		return nil, err
	}

	return dish, nil
}

func (service *dishService) UpdateDish(ctx context.Context, id int, data *dto.DishUpdation) error {
	if err := service.store.UpdateDish(ctx, id, data); err != nil {
		return err
	}

	return nil
}

func (service *dishService) DeleteDish(ctx context.Context, id int) error {
	if err := service.store.DeleteDish(ctx, id); err != nil {
		return err
	}

	return nil
}
