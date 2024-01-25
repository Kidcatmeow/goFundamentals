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
	b := 10
	for next := true; next; next = b < 5 {
		fmt.Println(b)
		b++
	}

	d := 0
	for next := true; next; next = d < 5 {
		fmt.Println(d)
		d++
	}

	//forever loop
	c := 0
	for {
		fmt.Println()
		c++
	}

	//Print from slices or array
	slice := []string{"Kidcat", "Mochi", "Honey", "Lemon"}
	// array := [4]string{"Kidcat", "Mochi", "Honey", "Lemon"}
	fmt.Println("\n for basic")

	for i := 0; i < len(slice); i++ {
		value := slice[i]
		fmt.Println(i, ":", value)
	}

	//range return two values = index and value
	fmt.Printf("\n for range slice")
	for index, value := range slice {
		fmt.Println(index, ":", value)
	}

	//but if we don't want to use index with range then put _ instead
	//range return two values = index and value
	fmt.Println("\nOnly value")
	for _, value := range slice {
		fmt.Println("Only value:", value)
	}

	//if we want only index
	fmt.Println("\nOnly index")
	for index := range slice {
		fmt.Println("Only index", index)
	}

}
