package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the guessing game.")

	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(20)

	max_attempt := 5
	var guess int
	var counter int

	for {
		if counter < max_attempt {
			fmt.Println("Please guess a number: ")
			fmt.Scan(&guess)

			if guess < secretNumber {
				fmt.Println("Too small")

			} else if guess > secretNumber {
				fmt.Println("Too big")

			} else {
				fmt.Println("Congrat!/n You guessed right.")
				fmt.Printf("You guess %d times", counter)
				break
			}
		} else {
			fmt.Println("You have exceeded your max attempts(5). Please try again")
			break
		}

		counter++
	}

}
