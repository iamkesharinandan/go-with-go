package main

import "os"

func main() {
	username, password := "knandan", "password"

	input := os.Args
	if len(input) <= 1 {
		println("Usage: [username] [password]")
	} else if len(os.Args) == 2 {
		println("Please enter the password")
	} else if os.Args[1] != username {
		println("Invalid Username")
	} else if os.Args[2] != password {
		println("Invalid Password")
	} else {
		println("You are logged in")
	}
}
