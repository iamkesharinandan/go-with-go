package main

import "fmt"

func main() {
	var brand = "Google"
	fmt.Printf("%q\n", brand) // q means quoted strings
	fmt.Printf("%s\n", brand) // s means strings

	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	// fmt.Printf("Planet: %v\n", planet)
	// fmt.Printf("Distance: %v millions kms\n", distance)
	// fmt.Printf("Orbital Period: %v days\n", orbital)
	// fmt.Printf("Does: %v has life? %v\n", planet, hasLife)

	// fmt.Printf("%v is %v away. Think! %[2]v kms! %[1]v OMG!\n", planet, distance)

	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d millions kms\n", distance)
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Does: %s has life? %t\n", planet, hasLife)

	fmt.Printf("%v is %v away. Think! %[2]v kms! %[1]v OMG!\n", planet, distance)

	// Change PrintF precision for float
	fmt.Println()
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
}

// d in %d means integer
// T in %T means type of the variable
