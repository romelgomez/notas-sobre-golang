package main

import (
	"fmt"
)

func main() {

	// 1 form
	var message string
	message = "Hello Word"

	// 2 form
	var message3 string = "Hello Word 3"

	// 3 form
	message2 := "Hello Word 2"

	// 1 form
	var a, b, c int

	// 2 form
	var a2, b2, c2 int = 1, 2, 3

	// 3 form
	a3, b3, c3 := 1, false, 3

	fmt.Println(message)
	fmt.Println(message2)
	fmt.Println(message3)
	fmt.Println(a, b, c)
	fmt.Println(a2, b2, c2)
	fmt.Println(a3, b3, c3)

}
