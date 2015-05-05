package main

import (
	"fmt"
)

type User struct {
	name     string
	lastName string
	pc 	PC
}

type PC  struct {
	box        string
	memory     string
	procesador string
}

type Users []User

func main() {

	// 0 form
	randomUser := User{
		"romel",
		"Gomez",
		PC{},
	}

	// 1 form
	randomUser1 := User{
		"romel",
		"Gomez",
		PC{
			"corsair 800D",
			"intel",
			"core 2duo",
		},
	}

	// 2 form
	var randomUser2 User
	randomUser2.name = "romel"
	randomUser2.lastName = "Gomez"
	randomUser2.pc.box = "corsair 800D"
	randomUser2.pc.memory = "intel"
	randomUser2.pc.procesador = "core 2duo"

	// 4 form
	randomUsers := Users{
		0:{
			"romel",
			"Gomez",
			PC{
				"corsair 800D",
				"intel",
				"core 2duo",
			},
		},
		1:{
			"Rudy",
			"Gomez",
			PC{},
		},
		2:{
			"Dilia",
			"Gomez",
			PC{
				"generic",
				"intel",
				"core 2duo",
			},
		},
	}

	fmt.Println(randomUser)
	fmt.Println(randomUser1)
	fmt.Println(randomUser2)
	fmt.Println(randomUsers)

}
