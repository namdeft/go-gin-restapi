package model

import (
	"time"
)

type Ingredient struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Quantity      int       `json:"quantity"`
	Import_Date   time.Time `json:"import_date"`
	Export_Date   time.Time `json:"export_date"`
	Counting_Unit int       `json:"counting_unit"`
}

func (Ingredient) TableName() string {
	return "ingredient"
}
