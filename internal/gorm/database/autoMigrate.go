package database

import (
	"golang/internal/gorm/model"
)

func AutoMigrate() {
	// 迁移 schema
	err := DB.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		panic(err)
	}
}
