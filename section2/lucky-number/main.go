package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5
	usage    = `Welcome to the Luck Number Game! 
The program will pick %d random numbers.
You mission is to guess one those numbers.

The greater your number is, the harder it gets.

Wanna play?`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 0; turn < 5; turn++ {
		n := rand.Intn(guess + 1)
		if n == guess {
			fmt.Println("You win.")
			return
		}
	}

	fmt.Println("You lost... try again")
}
