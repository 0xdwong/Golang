package main

import (
	"fmt"

	"gorm.io/driver/postgres"
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

func Select() {
	dsn := "host=localhost user=test password=123456 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	var user User

	// 根据整型主键查找
	if db.First(&user, 1).Error != nil {
		fmt.Println("Not found")
	} else {
		fmt.Println("Found user:", user.Name)
	}

}

func main() {
	Select()
}
