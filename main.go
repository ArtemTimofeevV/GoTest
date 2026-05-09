package main

import "fmt"

type User struct {
	Name   string
	Rating float64
}

func (u User) ChangeUser() {
	fmt.Print("Всем привет! Меня зовут ", u.Name, ".\n")
	fmt.Println("Из-за моего рейтинга я не могу пройти тест.")
	fmt.Print("Мой текущий рейтинг: ", u.Rating, ".\n")
	fmt.Print("Я хочу поменять имя на Егор.\n")
	u.Name = "Егор"
	fmt.Print("Теперь меня зовут ", u.Name, ".\n")
}

func (u *User) RatingUp(rating float64) {
	if u.Rating+rating <= 10 {
		u.Rating += rating
		fmt.Printf("Рейтинг повышён до ")
	} else {
		fmt.Printf("Я не смог поднять рейтинг\n")
	}
}

func main() {
	User := User{
		Name:   "Александр",
		Rating: 1.0,
	}
	User.ChangeUser()
	User.RatingUp(5.1)
	fmt.Println(User.Rating)
}
