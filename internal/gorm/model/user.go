package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  uint
}

func (User) TableName() string {
	return "user"
}
