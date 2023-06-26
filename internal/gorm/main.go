package main

import (
	"golang/internal/gorm/database"
	"golang/internal/gorm/demo"
)

func main() {
	database.InitPgsql()
	demo.Run()
}
