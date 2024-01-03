package business

import (
	"context"
	"gin-restapi/internal/dish/model"
)

type CreateDishStorage interface {
	CreateDish(ctx context.Context, data *model.DishCreation) error
}

type createDishBusiness struct {
	store CreateDishStorage
}

func CreateDishBusiness(store CreateDishStorage) *createDishBusiness {
	return &createDishBusiness{store: store}
}

func (business *createDishBusiness) CreateNewDish(ctx context.Context, data *model.DishCreation) error {
	if err := business.store.CreateDish(ctx, data); err != nil {
		return err
	}

	return nil
}
