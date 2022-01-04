package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		n   int
		err error
	)
	// Short-If
	if a := os.Args; len(a) != 2 {

		// Only a variable
		fmt.Println("Give me a number.")
	} else if n, err = strconv.Atoi(a[1]); err != nil {

		// only a, n and err variable
		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		n *= 2
		// all the variable in the if statement
		fmt.Printf("%s * 2 %d\n", a[1], n*2)
	}

	fmt.Printf("n is %d. Oooo - you've been shadowed ;) ", n)
}
