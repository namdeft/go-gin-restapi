package storage

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

func SQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
