package main

import (
	"fmt"
)

func main() {

	//	boolean
	var b bool = true
	fmt.Printf("boolean: %t\n", b)

	// string
	var s string = "Hello Word"
	fmt.Printf("string: %s\n", s)

	// pointer
	p := int(1)
	fmt.Printf("%p\n", &p)

	// custom type
	type customType string
	var randomVar customType = "Type Example"
	fmt.Println(randomVar)

}

// https://golang.org/pkg/fmt/

//basic types
//	bool
//	string
//	int, int8, int16, int32, int64
//	uint,uint8, uint16, uint32, uint64, uintptr
//	byte (uint8)
//	rune (int32), like a char
//	float32,float64
//	complex64,complex128
//Other types
//	array
//	slice
//	struct
//	pointer
//	function
//	interface
//	map
//	channel

// Use float64 whenever possible, all the math package expect that type.
