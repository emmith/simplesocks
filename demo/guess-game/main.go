package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const MAX_NUM = 100

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	secretNum := rand.Intn(MAX_NUM)
	fmt.Printf("The secret number is %v\n", secretNum)

	fmt.Printf("Please input your guess\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occured while read input, please try again %v\n", err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")
		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Printf("Invalid input %v, please input a integer", input)
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
