package main

import (
	"fmt"
	"math"
)

func main() {
	var x, maxi, mini, med float64
	fmt.Scan(&x)
	maxi = math.Max(math.Max(float64((int(x)%10)), float64(int(x)/100)), float64(int(x)/10%10))
	mini = math.Min(math.Min(float64((int(x)%10)), float64(int(x)/100)), float64(int(x)/10%10))
	fmt.Println(maxi, mini)
	if float64(int(x)/100) != maxi && float64(int(x)/100) != mini {
		med = float64(int(x) / 100)
		fmt.Println(1, "aaa")
	}
	if float64(int(x)%10) != maxi && float64(int(x)%10) != mini {
		med = float64(int(x) % 10)
		fmt.Println(2, "bbb")
	}
	if float64(int(x)/10%10) != maxi && float64(int(x)/10%10) != mini {
		med = float64(int(x) / 10 % 10)
		fmt.Println(3, "ccc")
	}
	fmt.Println(med)
	if maxi-mini == med {
		fmt.Println("Число интересное")
	} else {
		fmt.Print("Число неинтересное")
	}
}
