package lib

import (
	"fmt"
	"time"
)

// if-else case
func If() {
	num := 33

	if num%2 == 0 {
		fmt.Println("Even number")
	} else if num == 15 {
		fmt.Println("The number is 15")
	} else {
		fmt.Println("Odd number")
	}

}

// switch case
// switch case in go automatically break when it result goes in each case
// use 'fallthrough' if you want it to run through all cases below except default case
func Switch() {

	today := time.Tuesday

	switch today {
	case time.Monday:
		fmt.Println("This is Monday")
	case time.Tuesday:
		fmt.Println("This is Tuesday")
		fallthrough
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Printf("Today is %v\n", today)
	}

}
