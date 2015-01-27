package test

import (
	"fmt"
)

type User struct {
	Name string
	LastName string
	Age	string
}


func RunLoop (data []User){

		for _, v := range data{
			fmt.Println(v.Name + " " +  v.LastName + " " + v.Age)
		}

}


type Status func(string) ()

func Hi(name, lastName, id string) (salute string, youId string)  {
	salute = " Hello " + name + " " + lastName
	youId =	" Your id is: " + id
	return
}

func SayHi(user []User){

	for _, v := range user{

		prefix := getPrefix(v.Name)

		salute := " Hello " + prefix + v.Name + " " + v.LastName + ", Your Age is?: " + v.Age

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

func TypeSwitchTest (x interface{}){

	switch x.(type) {
		case int:
			fmt.Println("is int")
		case string:
			fmt.Println("is string")
		default:
			fmt.Println("is unknown")
	}

}



func Success( text string) {
	message := "Success: " + text
	fmt.Println(message)
}

func Error( text string) {
	message := "Error: " + text
	fmt.Println(message)
}


func Message(message string) Status {

	return func(s string){
		fmt.Println(s + " " + message)
	}

}
