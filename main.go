package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	if rand.Intn(5) == 3 {
		fmt.Println("Победа")
	} else {
		fmt.Print("Проигрыш!")
	}
}
