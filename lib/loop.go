package main

import "fmt"

func main() {

	//for-loop
	//condition checked before doing loop body
	//execute the loop body  for specific numeber of time (number of iterations known)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//while-loop
	//condition checked before doing loop body
	//execute the loop body until the condition is false
	a := 0
	for a < 10 {
		fmt.Println(a)
		a++
	}

	//do-while loop
	//loop body execute atleast once
	b := 0
	for b := true; b; b < 10 {
		fmt.Println(b)
		b++
	}

	// //forever loop
	c := 0
	for {
		fmt.Println()
		c++
	}
}
