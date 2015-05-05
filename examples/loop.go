package main

import (
	"fmt"
)

type User struct {
	Name     string
	LastName string
}

type Users []User

func main() {

	users := Users{
		0: {
			"Romel Javier",
			"Gomez Herrera",
		},
		1: {
			"Rudy Alberto",
			"Gomez Herrera",
		},
		2: {
			"Dilia",
			"Gomez Herrera",
		},
	}

	sayHi(users)

}



func sayHi(users Users){

	for _, v := range users{

		prefix := getPrefix(v.Name)

		salute := " Hello " + prefix + v.Name + " " + v.LastName

		fmt.Println(salute)

	}

}

func getPrefix( name string) ( prefix string) {

	switch name {
		case "Romel":
			prefix = "Mr, "
		case "Dilia":
			prefix = "Mrs, "
		default:
			prefix = "Dude, "
	}

	return
}

