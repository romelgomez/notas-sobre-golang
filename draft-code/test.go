package main

import (
	"./test"
)


func main(){

		users := test.Users{
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

	  users.Add()

//	var slice2 []int = make([]int, 3, 4)
//
//	slice2[0] = 1
//	slice2[1] = 10
//	slice2[2] = 12
//
//
//	users := map[string] test.User{
//		"1":{
//			"Romel Javier",
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

}
