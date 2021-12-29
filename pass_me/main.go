package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOk = "Access granted to %q.\n"

	user     = "jack"
	password = "1888"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
	}

	u, p := args[1], args[2]
	if u != user {
		fmt.Printf(errUser, u)
	} else if p != password {
		fmt.Printf(errPwd, u)
	} else {
		fmt.Printf(accessOk, u)
	}
}
