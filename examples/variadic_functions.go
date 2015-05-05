package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Variadic Functions; 0:Hi, 1:Hello, 2:Hola
	v := variadicFunction(0, "Hi", "Hello", "Hola")
	fmt.Println(v + " Word")
}

func variadicFunction(i int, v ...string) string {
	var vLength int = len(v)
	var vLength64 int64 = int64(vLength)
	fmt.Println("'v' var has length of: " + strconv.FormatInt(vLength64, 10))
	fmt.Println(v)
	return v[i]
}
