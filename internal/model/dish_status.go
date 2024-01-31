package model

import (
	"errors"
	"fmt"
)

type DishStatus int

const (
	DishStatusUnavailable DishStatus = iota
	DishStatusAvailable
	DishStatusDeleted
)

var allDishStatuses = [3]string{
	"unavailable",
	"available",
	"deleted",
}

func (dish DishStatus) String() string {
	return allDishStatuses[dish]
}

func parseStrToDishStatus(s string) DishStatus {
	for i := range allDishStatuses {
		if allDishStatuses[i] == s {
			return DishStatus(i)
		}
	}

	return DishStatusUnavailable
}

func (dish *DishStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("fail to scan data from sql1: %s", value))
	}

	v := parseStrToDishStatus(string(bytes))

	*dish = v

	return nil
}

func (dish *DishStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", dish.String())), nil
}
