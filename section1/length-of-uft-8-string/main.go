package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "Késhari"      // has 7 chars
	fmt.Println(len(name)) // returns 8 - has error

	fmt.Println(utf8.RuneCountInString(name)) // retuens 7  - correct
}
