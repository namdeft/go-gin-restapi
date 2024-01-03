package storage

import (
	"context"
	"gin-restapi/internal/user/model"
)

func (s *sqlStore) Register(ctx context.Context, data *model.RegisterInput) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
