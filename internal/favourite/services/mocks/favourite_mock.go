package mocks

import (
	"context"
	"gin-restapi/internal/favourite/model"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) AddFavourite(ctx context.Context, userId int, dishId int) error {
	args := mock.Called(ctx, userId, dishId)
	return args.Error(1)
}

func (mock *MockRepository) GetFavourites(ctx context.Context, userId int) ([]model.Dish, error) {
	args := mock.Called(ctx, userId)
	result := args.Get(0)
	return result.([]model.Dish), args.Error(1)
}

func (mock *MockRepository) DeleteFavourite(ctx context.Context, userId int, dishId int) error {
	args := mock.Called(ctx, userId, dishId)
	return args.Error(1)
}
