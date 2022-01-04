package main

// // version := 0 // Not recomended and not valid
// var version int // Recomended and Valid

// func main() {
// 	// score := 0 // Do no use shot declaration when you are unsure about the value
// 	var score int // Recomended and Valid and default value is set 0 automatically

// 	// Variable groupings
// 	var (
// 		// Related
// 		video string

// 		// Closely Related
// 		duration int //
// 		current  int //
// 	)
// }

func main() {
	// var width, height = 800, 600 // Don't do this

	// Instead do this -
	width, height := 800, 600

	// // Don't do this
	// width = 10 // assign 50 the width
	// color := "red"
}
width, color = 100, "red"
