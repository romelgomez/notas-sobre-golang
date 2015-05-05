package main

import (
	"fmt"
)

func main() {

}

func TypeSwitchTest(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println("is int")
	case string:
		fmt.Println("is string")
	default:
		fmt.Println("is unknown")
	}

}
