package main

import (
	"mygorm/database"
	"mygorm/demo"
)

func main() {
	database.InitPgsql()
	demo.Run()
}
