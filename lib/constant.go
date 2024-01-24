package lib

import "fmt"

func Const() {

	// var variable values can be altered
	var day string = "Monday"
	fmt.Println("day:", day)

	day = "Kidcat's day"
	fmt.Println("day:", day)

	//const variables cannot be altered
	const dessert string = "Black Forest"
	fmt.Println("Dessert:", dessert)
	//Changing const variables results in error
	// dessert = "Cherry pie"

	//To declare multiple constant variables
	const (
		Sunday    = 0
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
		Thursday  = 4
		Friday    = 5
	)

	fmt.Println("Sunday :", Sunday)
	fmt.Println("Monday :", Monday)
	fmt.Println("Tuesday :", Tuesday)
	fmt.Println("Wednesday :", Wednesday)
	fmt.Println("Thursday :", Thursday)
	fmt.Println("Friday :", Friday)

	//Manually type 0-5 can be exhausting so go has 'iota' to auto increment the values
	const (
		Zero = iota
		One
		Two
		Three
		Four
		Five
	)

	fmt.Println("Zero :", Zero)
	fmt.Println("One :", One)
	fmt.Println("Two :", Two)
	fmt.Println("Three :", Three)
	fmt.Println("Four :", Four)
	fmt.Println("Five :", Five)

}
