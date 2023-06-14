package database

import (
	"mygorm/model"
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
