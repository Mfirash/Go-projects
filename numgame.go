package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var guesse string
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100) + 1
	fmt.Println("guess the number from 1 to 100")
	for {
		fmt.Scanln(&guesse)

		guess, err := strconv.Atoi(guesse)

		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		if guess < num {
			fmt.Println("too low")
		} else if guess > num {
			fmt.Println("too high")
		} else {
			fmt.Println("you found it it was:", num)
			break
		}
	}
}
