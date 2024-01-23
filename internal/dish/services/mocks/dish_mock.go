package mocks

import (
	"context"
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/dto"
	"gin-restapi/internal/dish/model"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) CreateDish(ctx context.Context, data *dto.DishCreation) error {
	args := mock.Called()
	return args.Error(1)
}

func (mock *MockRepository) GetDishes(ctx context.Context, paging *common.Paging) ([]model.Dish, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.Dish), args.Error(1)
}

func (mock *MockRepository) GetDish(ctx context.Context, id int) (*model.Dish, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*model.Dish), args.Error(1)
}

func (mock *MockRepository) UpdateDish(ctx context.Context, id int, data *dto.DishUpdation) error {
	args := mock.Called()
	return args.Error(1)
}

func (mock *MockRepository) DeleteDish(ctx context.Context, id int) error {
	args := mock.Called()
	return args.Error(1)
}
