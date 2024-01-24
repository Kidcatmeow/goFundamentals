package lib

import "fmt"

/*
There are public and private function
if you want the function to be public then start the naming with capital letter
*/

// Function that receive parameters as value
func Info(name, msg string, age int) {
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("My message is %s\n", msg)
	fmt.Printf("I am %d years old\n", age)
}

// Function that return value
// function todayonly  returns string -->  มันให้ค่ากลับมาอย่างเดียวเพราะฉะนั่น ต้องประกาศตัวแปรมารับค่า
// so it cannot print by its own , you need to store the return value inside a variable before printing
func Today() string {
	return "Eat this"
}

// function that receives and return 2 values
func Swap(x, y string) (string, string) {
	return y, x
}

// function เป็น variable
/*
The function below is the same as but it turned it into variable
func Plus(a int, b int) int {
	return a + b
}
*/
var Plus = func(a int, b int) int { return a + b }

// This means that a function is merely like a variable in go
// so a function can be send anywhere like a mere variable can
// and a function can be a parameter for another function
func Cal(op func(int, int) int) {
	r := op(4, 5)
	fmt.Printf("result :%v\n", r)
}

var Minus = func(a int, b int) int { return a - b }
