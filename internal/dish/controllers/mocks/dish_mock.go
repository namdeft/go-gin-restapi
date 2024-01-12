package mocks

import (
	"context"
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/dto"
	"gin-restapi/internal/dish/model"

	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetDishes(ctx context.Context, paging *common.Paging) ([]model.Dish, error) {
	args := m.Called(ctx, paging)
	return args.Get(0).([]model.Dish), args.Error(1)
}

func (m *MockService) GetDish(ctx context.Context, id int) (*model.Dish, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Dish), args.Error(1)
}

func (m *MockService) CreateNewDish(ctx context.Context, data *dto.DishCreation) error {
	args := m.Called()
	return args.Error(1)
}

func (m *MockService) UpdateDish(ctx context.Context, id int, data *dto.DishUpdation) error {
	args := m.Called(ctx, id, data)
	return args.Error(1)
}

func (m *MockService) DeleteDish(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(1)
}
