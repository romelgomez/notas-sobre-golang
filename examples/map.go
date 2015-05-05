package main

import (
	"fmt"
)

type User struct {
	Name     string
	LastName string
	Age      string
}

type Users []User

func main() {

	users := Users{
		0: {
			"Romel Javier",
			"Gomez Herrera",
			"30",
		},
		1: {
			"Rudy Alberto",
			"Gomez Herrera",
			"25",
		},
		2: {
			"Dilia",
			"Gomez Herrera",
			"24",
		},
	}

	fmt.Println(users)

}
