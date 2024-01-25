package main

import "fmt"

type User struct {
	Username  string
	Firstname string
	Lastname  string
}

func info(u User) {
	fmt.Printf("Username : %v\n", u.Username)
	fmt.Printf("Firstname : %v\n", u.Firstname)
	fmt.Printf("Lastname : %v\n", u.Lastname)
}

// convert function info to a method of User
func (u User) info() {
	fmt.Printf("Username : %v\n", u.Username)
	fmt.Printf("Firstname : %v\n", u.Firstname)
	fmt.Printf("Lastname : %v\n", u.Lastname)
}

func main() {

	//Initialize struct
	//insert values into the struct

	u := User{
		Username:  "Kidcat",
		Firstname: "Watsamon",
		Lastname:  "Phongwanit",
	}

	info(u)

	//call method of user
	u.info()

}
