package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MaxNum = 100

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	secretNum := rand.Intn(MaxNum)
	fmt.Printf("The secret number is %v\n", secretNum)

	fmt.Printf("Please input your guess\n")

	var guess int

	for {
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Printf("An error occured while read input, please try again %v\n", err)
			continue
		}

		if guess == secretNum {
			fmt.Println("Guess right!")
			break
		} else if guess < secretNum {
			fmt.Println("It is smaller than secret number")
		} else {
			fmt.Println("It is bigger than secret number")
		}
	}

}
