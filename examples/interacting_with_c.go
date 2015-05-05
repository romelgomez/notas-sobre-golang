package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
)

func main() {
	fmt.Println("random int front C:", int(C.random()))
}
