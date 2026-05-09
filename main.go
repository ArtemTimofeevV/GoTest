package main

import "fmt"

type user struct {
	name    string
	age     int
	number  string
	isclose bool
	rating  float64
}

func main() {
	user := user{
		name:    "Сергей",
		age:     20,
		number:  "+7 (985) 234 - 43 - 41",
		isclose: true,
		rating:  0.1,
	}

	fmt.Println("Users: ", user)
	fmt.Println("Age: ", user.age)
}
