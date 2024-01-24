package lib

import "fmt"

func Different_print() {
	//Print newline
	fmt.Println("Hello Kidcat")

	//Print format
	fmt.Printf("Hello %s!!!\n", "Kidcat")
	fmt.Println("Next line")

	fmt.Printf("Hello %s!!!", "Kidcat")
	fmt.Println("Next line")

	//Print numbers
	fmt.Printf("Hello %d!!!\n", 66)

	//Print anything
	fmt.Printf("Hello %v!!!\n", 66)
	fmt.Printf("Hello %v!!!\n", "Gopher")

	//Print many things at once
	fmt.Printf("hello %v %v!!!!\n", "Gopher", 66)

	//Print with a quote
	fmt.Printf("%q\n", name_packageScope)
}
