package main

import (
	"fmt"
)

func main() {
	// Variadic Functions; 0:Hi, 1:Hello, 2:Hola
	v := variadicFunction(0, "Hi", "Hello", "Hola")
	fmt.Println(v + " Word")
}

func variadicFunction(i int, v ...string) string {
	fmt.Printf("'v' var has length of: %d \n", len(v))
	fmt.Println(v)
	return v[i]
}
