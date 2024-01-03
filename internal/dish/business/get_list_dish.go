package business

import (
	"context"
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/model"
)

type GetDishesStorage interface {
	GetDishes(ctx context.Context, paging *common.Paging) ([]model.Dish, error)
}

type getDishesBusiness struct {
	store GetDishesStorage
}

func GetDishesBusiness(store GetDishesStorage) *getDishesBusiness {
	return &getDishesBusiness{store: store}
}

func (business *getDishesBusiness) GetDishes(ctx context.Context, paging *common.Paging) ([]model.Dish, error) {
	dishes, err := business.store.GetDishes(ctx, paging)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}
