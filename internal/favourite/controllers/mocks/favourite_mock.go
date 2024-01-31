package mocks

import (
	"context"
	"gin-restapi/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (mock *MockService) AddFavourite(ctx context.Context, userId int, dishId int) error {
	args := mock.Called(ctx, userId, dishId)
	return args.Error(0)
}

func (mock *MockService) GetFavourites(ctx context.Context, userId int) ([]model.Dish, error) {
	args := mock.Called(ctx, userId)
	result := args.Get(0)
	return result.([]model.Dish), args.Error(1)
}

func (mock *MockService) DeleteFavourite(ctx context.Context, userId int, dishId int) error {
	args := mock.Called(ctx, userId, dishId)
	return args.Error(0)
}

func (mock *MockService) CheckDishExists(dishId int) error {
	args := mock.Called(dishId)
	return args.Error(0)
}
