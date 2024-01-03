package storage

import (
	"context"
	"gin-restapi/internal/dish/model"
	"time"
)

func (s *sqlStore) DeleteDish(ctx context.Context, id int) error {
	if err := s.db.Table(model.Dish{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"status":     "deleted",
		"deleted_at": time.Now(),
	}).Error; err != nil {
		return err
	}

	return nil
}
