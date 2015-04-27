package main

import (
	"fmt"
)

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

func boolExample(){
	var bol = false
	fmt.Println(bol)
}

func stringExample(){

	// 1 form
	var message string
	message = "Hello Word"

	// 2 form
	var message3 string = "Hello Word 3"

	// 3 form
	message2 := "Hello Word 2"

	// 1 form
	var a, b, c int

	// 2 form
	var a2, b2, c2 int = 1, 2, 3

	// 3 form
	a3, b3, c3 := 1, false, 3

	fmt.Println(message)
	fmt.Println(message2)
	fmt.Println(message3)
	fmt.Println(a, b, c)
	fmt.Println(a2, b2, c2)
	fmt.Println(a3, b3, c3)

}

func intExample(){
//	var a,b,c int
//	var a,b,c int = 1,2,3
}

func uintExample() {

}

func byteExample() {

}

func runeExample() {

}

func floatExample() {

}

func complexExample() {

}

func typeExample() {
	type mockString string

	var randomVar mockString = "Type Example"

	fmt.Println(randomVar)
}


func arrayExample(){
	
}

func sliceExample(){
	
}

func structExample(){

	type User struct {
		name string
		lastName string
	}

	// 1 form
	randomUser := User{
		"romel",
		"Gomez",
	}

	// 2 form
	var randomUser2 User
	randomUser2.name 		= "javier"
	randomUser2.lastName	= "herrera"

	fmt.Println(randomUser)
	fmt.Println(randomUser2)

	// 3 form
	type User2 struct {
		name string
		lastName string
		pc struct {
			box string
			memory string
			procesador string
		}
	}

	var randomUser3 User2

	randomUser3.name = "romel"
	randomUser3.lastName = "Gomez"
	randomUser3.pc.box = "corsair 800D"
	randomUser3.pc.memory = "intel"
	randomUser3.pc.procesador = "core 2duo"


	fmt.Println(randomUser3)


}

func pointerExample(){
//	un tipo de variable especial que mantiene la dirección de otra variable en memoria
	hello := "Hello Word"
	pointer := &hello
	*pointer = "Hello Romel"

	fmt.Println(hello, pointer, *pointer)
}


type User struct {
	Name string
	LastName string
	Age	string
}

func returningMultipleValues(user User) (userData string, plusData string){
	userData = "Name: " + user.Name + ", LastName:" + user.LastName + ", Age:" + user.Age
	plusData = "This day is sunday"
	return
}

func functionExample(){

	var randomUser User

	randomUser.Name 		= "romel"
	randomUser.LastName 	= "javier"
	randomUser.Age 			= "30"

	printThis,_ := returningMultipleValues(randomUser)

	fmt.Println(printThis)



}
	
func interfaceExample(){
	
}

func mapExample() {

	type User struct {
		Name string
		LastName string
		Age	string
	}

	type Users []User

	users := Users{
		0:{
			"Romel Javier",
			"Gomez Herrera",
			"30",
		},
		1:{
			"Rudy Alberto",
			"Gomez Herrera",
			"25",
		},
		2:{
			"Dilia",
			"Gomez Herrera",
			"24",
		},
	}

	fmt.Println(users)


	//	var usersMap = map[int]User{}


//	func (users Users) Add (){
//		for k , v := range users {
//			usersMap[k] = v
//		}
//		fmt.Println(usersMap)
//	}
//
//
//	users.Add()


}

func channelExample() {

}

//	var slice2 []int = make([]int, 3, 4)
//
//	slice2[0] = 1
//	slice2[1] = 10
//	slice2[2] = 12
//
//
//	users := map[string] test.User{
//		"1":{
//			"Romel Javier",n
//			"Gomez Herrera",
//			"30",
//		},
//		"2":{
//			"Rudy Alberto",
//			"Gomez Herrera",
//			"25",
//		},
//	}
//
//	var user test.User
//	user.Name = "Dilia Alvanelys"
//	user.LastName = "Gomez Herrera"
//	user.Age = "24"
//
//	user.Add()
//
//	users["3"] = user
//
//	delete(users, "1")
//
//	if value, exist := users["1"]; exist {
//		fmt.Println(value)
//	} else {
//		fmt.Println("1 no existe")
//	}
//
//
//	for k , v := range users {
//		user := "[Id] " + k + ", [Name] " + v.Name + ", [LastName] " + v.LastName + ", [Age] " + v.Age + ";"
//		fmt.Println(user)
//	}
//
//	slice := []test.User{
//		{"Romel", "Gomez", "30"},
//		{"Rudy", "Gomez", "24"},
//		{"Dilia", "Gomez", "25"},
//	}
//
////	test.RunLoop(slice)
//
//	test.SayHi(slice);



func main(){

//	boolExample()
//	stringExample()
//	mapExample()
//	pointerExample()
//	typeExample()
//	structExample()
	functionExample()
}
