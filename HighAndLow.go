package main

import (
	"fmt"
	"math/rand"
        "time"
)

var hand string

func main(){
	rand.Seed(time.Now().UnixNano())
	fmt.Println("相手の数字", Dealer())
	fmt.Println("high and low?")
	fmt.Scan(&hand)

	yournumber := rand.Intn(13)

	switch hand {
	case "high":
		if Dealer() <= yournumber{
			fmt.Println("your number is", yournumber)
			fmt.Println("Win")
		} else {
			fmt.Println("your number is", yournumber)
			fmt.Println("Lose")
		}
	case "low":
		if Dealer() >= yournumber{
			fmt.Println("your number is", yournumber)
			fmt.Println("Win")
		} else {
			fmt.Println("your number is", yournumber)
			fmt.Println("Lose")
		}
		default: fmt.Println("what is that?")
	}
}

func Dealer() int {
	// 異なる乱数の生成
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(13)
}
