package main

type semester struct {
}

type speciality struct {
	name string
}

type courseProperties struct {
	academicLoad     string
	theoreticalHours string
	precedence       string
}

type course struct {
	Name       string
	Properties map[string]string
}

//	struct
// 		entity
// 			properties.

// course
// properties
// 	academic load, theoretical hours, precedence.
