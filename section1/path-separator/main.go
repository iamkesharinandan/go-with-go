package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string

	dir, file = path.Split("css/index.html")

	fmt.Println("Path: ", dir)
	fmt.Println("File: ", file)
}
