package lib

import "fmt"

func Pointer() {
	me := "Kidcat"
	fmt.Printf("You are %s\n", me)
	//use '&' to print memory location where 'me' variable was allocated
	// addr := &me
	//same as
	var addr *string = &me
	fmt.Println(addr)
	fmt.Printf("%T\n", addr)

	//now if we want to change the value of me,you can do it this way
	//go the the pointer of string 'me' and change the value tp 'penguin'
	*addr = "penguin"
	fmt.Printf("You are %s\n", me)

}
