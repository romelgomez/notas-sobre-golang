package main

type semester struct {
}

type speciality struct {
	Name string
}

type courseProperties struct {
	AcademicLoad     string
	TheoreticalHours string
	Precedence       string
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
