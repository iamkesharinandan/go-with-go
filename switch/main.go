package main

import (
	"fmt"
	"os"
)

func main() {
	city := os.Args[1]

	switch city {
	case "Paris", "Lyon":
		fmt.Println("France")
		break // Optional

	case "Tokyo":
		fmt.Println("Japan")
		break // Optional

	default:
		fmt.Println("Where?")
	}

	// Switch with Bool
	i := -0
	switch {
	case i > 0:
		fmt.Println("Positive")
	case i < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("zero")
	}
}
