package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("USER: ", os.Getenv("USER"))
	fmt.Println("HOME: ", os.Getenv("HOME"))
	fmt.Println("GOROOT: ", os.Getenv("GOROOT"))
	fmt.Println("GOPATH: ", os.Getenv("GOPATH"))
	fmt.Println("GOOS: ", os.Getenv("GOOS"))

}
