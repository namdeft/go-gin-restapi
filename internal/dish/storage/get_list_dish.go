package storage

import (
	"context"
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/model"
)

func (s *sqlStore) GetDishes(ctx context.Context, paging *common.Paging) ([]model.Dish, error) {
	offset := (paging.Page - 1) * paging.Limit

	var dishes []model.Dish
	if err := s.db.
		Table(model.Dish{}.TableName()).
		Count(&paging.Total).
		Offset(offset).
		Limit(paging.Limit).
		Find(&dishes).Error; err != nil {
		return nil, err
	}

	return dishes, nil
}
