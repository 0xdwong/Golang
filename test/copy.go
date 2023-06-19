package main

import (
	"fmt"
)

func Copy() {
	type Person struct {
		Name string
		Age  int
	}

	p1 := Person{"xxx", 18}
	// p2 := p1
	p2 := &p1

	p1.Age = 20
	fmt.Println(p2)
}
