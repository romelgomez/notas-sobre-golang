package test

type Semester struct {
}

type Speciality struct {
	Name string
}

type CourseProperties struct {
	AcademicLoad string
	TheoreticalHours string
	Precedence string
}

type Course struct {
	Name string
	Properties map[string] string;
}

//	struct
// 		entity
// 			properties.

// course
// properties
// 	academic load, theoretical hours, precedence.
