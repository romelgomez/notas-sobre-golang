package main

import (
	"fmt"
)

func main() {

//	struct
// 		entity
// 			properties.


// course
// properties
// 	academic load, theoretical hours, precedence.


type semester struct {

}

type speciality struct {
	name string
}

type courseProperties struct {
}

type course struct {
	name string
	properties map[string] courseProperties;
}



//p := new(courseProperties)
f := new(course)

f.name =	"Fisica"
//f.properties["academicLoad"] = p.academicLoad

fmt.Println(f)

//var m = map[string] string {"evga":"890","nvdia":"980"}

//fmt.Println(m["evga"]);

}
