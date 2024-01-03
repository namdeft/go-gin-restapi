package business

import (
	"context"
	"gin-restapi/internal/dish/model"
)

type UpdateDishStorage interface {
	UpdateDish(ctx context.Context, id int, data *model.DishUpdation) error
}

type updateDishBusiness struct {
	store UpdateDishStorage
}

func UpdateDishBusiness(store UpdateDishStorage) *updateDishBusiness {
	return &updateDishBusiness{store: store}
}

func (business *updateDishBusiness) UpdateDish(ctx context.Context, id int, data *model.DishUpdation) error {
	if err := business.store.UpdateDish(ctx, id, data); err != nil {
		return err
	}

	return nil
}
