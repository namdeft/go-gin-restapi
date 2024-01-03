package business

import (
	"context"
)

type DeleteDishStorage interface {
	DeleteDish(ctx context.Context, id int) error
}

type deleteDishBusiness struct {
	store DeleteDishStorage
}

func DeleteDishBusiness(store DeleteDishStorage) *deleteDishBusiness {
	return &deleteDishBusiness{store: store}
}

func (business *deleteDishBusiness) DeleteDish(ctx context.Context, id int) error {
	if err := business.store.DeleteDish(ctx, id); err != nil {
		return err
	}

	return nil
}
