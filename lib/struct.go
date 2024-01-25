package main

import "fmt"

type User struct {
	Username  string
	Firstname string
	Lastname  string
}

func main() {

	username := "Kidcat"
	firstname := "Watsamon"
	lastname := "Phongwanit"
	fmt.Println(username, firstname, lastname)

	//Initialize struct
	u := User{}
	//insert values into the struct
	u.Username = username
	u.Firstname = firstname
	u.Lastname = lastname
	//another wasy is
	v := User{
		Username:  username,
		Firstname: firstname,
		Lastname:  lastname,
	}

	//type of struct also include the package
	fmt.Printf("%#v\n", u)
	fmt.Printf("%#v\n", v)

	//print value of struct
	name := u.Username
	fmt.Println(name)
	//or just
	fmt.Println(u.Firstname)
}
