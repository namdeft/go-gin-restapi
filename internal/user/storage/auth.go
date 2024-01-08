package storage

import (
	"context"
	"gin-restapi/internal/user/dto"
	"gin-restapi/internal/user/model"
)

func (s *SqlStore) Login(ctx context.Context, data *dto.LoginInput) (*model.User, error) {
	var u model.User
	if err := s.db.Where("email = ?", data.Email).Find(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *SqlStore) Register(ctx context.Context, data *dto.RegisterInput) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
