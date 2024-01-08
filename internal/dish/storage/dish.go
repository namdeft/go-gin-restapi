package storage

import (
	"context"
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/dto"
	"gin-restapi/internal/dish/model"
	"time"
)

func (s *SqlStore) DeleteDish(ctx context.Context, id int) error {
	if err := s.db.Table(model.Dish{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"status":     "deleted",
		"deleted_at": time.Now(),
	}).Error; err != nil {
		return err
	}

	return nil
}

func (s *SqlStore) GetDish(ctx context.Context, id int) (*model.Dish, error) {
	var dish model.Dish
	if err := s.db.Where("id = ?", id).First(&dish).Error; err != nil {
		return nil, err
	}

	return &dish, nil
}

func (s *SqlStore) GetDishes(ctx context.Context, paging *common.Paging) ([]model.Dish, error) {
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

func (s *SqlStore) CreateDish(ctx context.Context, data *dto.DishCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *SqlStore) UpdateDish(ctx context.Context, id int, data *dto.DishUpdation) error {
	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}
