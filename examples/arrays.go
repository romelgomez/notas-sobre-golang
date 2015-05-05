package main

import "fmt"

func main() {
//https://tour.golang.org/moretypes/6

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}
