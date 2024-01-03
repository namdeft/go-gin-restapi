package business

import (
	"context"
	"gin-restapi/internal/dish/model"
)

type GetDishStorage interface {
	GetDish(ctx context.Context, id int) (*model.Dish, error)
}

type getDishBusiness struct {
	store GetDishStorage
}

func GetDishBusiness(store GetDishStorage) *getDishBusiness {
	return &getDishBusiness{store: store}
}

func (business *getDishBusiness) GetDish(ctx context.Context, id int) (*model.Dish, error) {
	dish, err := business.store.GetDish(ctx, id)
	if err != nil {
		return nil, err
	}

	return dish, nil
}
