package main

import (
	"./test"

)


func main(){

	slice := []test.User{
		{"Romel", "Gomez", "30"},
		{"Rudy", "Gomez", "24"},
		{"Dilia", "Gomez", "25"},
	}

//	test.RunLoop(slice)

	test.SayHi(slice);


}
