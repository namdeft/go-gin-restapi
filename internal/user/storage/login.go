package storage

import (
	"context"
	"gin-restapi/internal/user/model"
)

func (s *sqlStore) Login(ctx context.Context, data *model.LoginInput) (*model.User, error) {
	var u model.User
	if err := s.db.Where("email = ?", data.Email).Find(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}
