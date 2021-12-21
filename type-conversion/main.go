package main

import "fmt"

func main() {
	speed := 100 //speed is int
	force := 2.5 // foce is float64

	// speed = speed * force // throws an error - invalid operation: speed * force (mismatched types int and float64)

	// solution
	speed = speed * int(force)
	fmt.Println(speed)
}
