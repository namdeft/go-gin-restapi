package mocks

import (
	"gin-restapi/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) AddFavourite(userId int, dishId int) error {
	args := mock.Called(userId, dishId)
	return args.Error(1)
}

func (mock *MockRepository) GetFavourites(userId int) ([]model.Dish, error) {
	args := mock.Called(userId)
	result := args.Get(0)
	return result.([]model.Dish), args.Error(1)
}

func (mock *MockRepository) DeleteFavourite(userId int, dishId int) error {
	args := mock.Called(userId, dishId)
	return args.Error(1)
}

func (mock *MockRepository) CheckDishExists(dishId int) error {
	args := mock.Called(dishId)
	return args.Error(0)
}
