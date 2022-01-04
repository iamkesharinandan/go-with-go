package main

import "fmt"

func main() {
	speed := 100 //speed is int
	force := 2.5 // foce is float64

	// speed = speed * force // throws an error - invalid operation: speed * force (mismatched types int and float64)

	// Attempt 1
	//speed = speed * int(force)
	//fmt.Println(speed) // result is diffeent than what we expect

	// Attempt 2
	//speed = float64(speed) * force // thows an error - cannot use float64(speed) * force (type float64) as type int in assignment

	// Attempt 3
	speed = int(float64(speed) * force)
	fmt.Println(speed)
}
