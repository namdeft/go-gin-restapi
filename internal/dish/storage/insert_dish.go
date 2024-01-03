package storage

import (
	"context"
	"gin-restapi/internal/dish/model"
)

func (s *sqlStore) CreateDish(ctx context.Context, data *model.DishCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
