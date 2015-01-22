package main

import (
	"fmt"
)

type User struct {
	name string
	lastName string
	id	string
}

type Semester struct {

}

type Speciality struct {
	name string
}

type CourseProperties struct {
	academicLoad string
	theoreticalHours string
	precedence string
}

type Course struct {
	name string
	properties map[string] string;
}

type Status func(string) ()

func Hi(name, lastName, id string) (salute string, youId string)  {
	salute = " Hello " + name + " " + lastName
	youId =	" Your id is: " + id
	return
}

func sayHi(user User, do Status){

	name, id := Hi(user.name, user.lastName, user.id)

	do(name);
	do(id);

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


func main() {


	var user = User{}
	user.name 		= "romel"
	user.lastName = "gomez"
	user.id 			= "16"

	sayHi(user,Message("!!!"));

//var a = CourseProperties{}
//a.academicLoad = "1"
//a.precedence	= "0"
//a.theoreticalHours = "10"

//fmt.Println(a)

}


//	struct
// 		entity
// 			properties.

// course
// properties
// 	academic load, theoretical hours, precedence.
