package main


type User struct {
	Name     string
	LastName string
	Password string
}

func main(){
}

func Hi(name, lastName, id string) (salute string, youId string)  {
	salute = " Hello " + name + " " + lastName
	youId =	" Your id is: " + id
	return
}
