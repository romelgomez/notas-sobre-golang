package main

import (
	"fmt"
)

type User struct {
	Name     string
	LastName string
	Password string
}

func main() {
	saludo, _ := Hi("romel", "gomez", "1")
	fmt.Printf("Mi nombre es %s y mi id es desconocido", saludo)
}

func Hi(name, lastName, id string) (salute string, youId string) {
	salute = " Hello " + name + " " + lastName
	youId = " Your id is: " + id
	return
}
