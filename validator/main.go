package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name  string `validate:"required"`
	Age   uint8  `validate:"gte=0,lte=130"`
	Email string `validate:"required,email"`
}

func main() {
	user1 := User{"username", 18, "xxx@gmail.com"}
	user2 := User{"", 18, "xxx@gmail.com"}
	user3 := User{"username", 200, "xxx@gmail.com"}
	user4 := User{"username", 18, "xxxx"}

	var users = []User{user1, user2, user3, user4}
	validateUsers(users)
}

func validateUsers(users []User) {
	validate := validator.New()

	for _, user := range users {
		if err := validate.Struct(user); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("validation pass")
		}
	}
}
