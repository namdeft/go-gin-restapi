package storage

import (
	"context"
	"gin-restapi/internal/dish/model"
)

func (s *sqlStore) GetDish(ctx context.Context, id int) (*model.Dish, error) {
	var dish model.Dish
	if err := s.db.Where("id = ?", id).First(&dish).Error; err != nil {
		return nil, err
	}

	return &dish, nil
}
