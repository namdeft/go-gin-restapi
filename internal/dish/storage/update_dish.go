package storage

import (
	"context"
	"gin-restapi/internal/dish/model"
)

func (s *sqlStore) UpdateDish(ctx context.Context, id int, data *model.DishUpdation) error {
	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}
