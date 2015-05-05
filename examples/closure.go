package main

import (
	"fmt"
)

type Print func(string)

func main() {
	functionTypes(closureExample("..."))
}

func functionTypes(do Print) {
	str := "Closure Example"
	do(str)
}

func closureExample(plus string) Print {
	return func(s string) {
		fmt.Println(s + " " + plus)
	}
}
