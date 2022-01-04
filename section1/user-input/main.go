package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path: ", os.Args[0])
	fmt.Println("1st Arg: ", os.Args[1])
	fmt.Println("2nd Arg: ", os.Args[2])
	fmt.Println("3rd Arg: ", os.Args[3])

	fmt.Println("No of item insite os.Args", len(os.Args))
}
