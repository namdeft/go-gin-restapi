package repository

import (
	"context"
	"gin-restapi/internal/user/dto"
	"gin-restapi/internal/user/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(ctx context.Context, data *dto.RegisterInput) error
	Login(ctx context.Context, data *dto.LoginInput) (*model.User, error)
}

type userConnection struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		db: db,
	}
}

func (s *userConnection) Login(ctx context.Context, data *dto.LoginInput) (*model.User, error) {
	var u model.User
	if err := s.db.Table(model.User{}.TableName()).Where("email = ?", data.Email).Find(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *userConnection) Register(ctx context.Context, data *dto.RegisterInput) error {
	if err := s.db.Table(model.User{}.TableName()).Create(&data).Error; err != nil {
		return err
	}

	return nil
}
