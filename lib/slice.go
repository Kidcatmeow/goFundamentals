package lib

import "fmt"

func Slice() {
	langs := []string{"golang", "python", "java"}
	fmt.Printf("langs: %#v\n", langs)

	//to print certain index in slice
	n := langs[0]
	fmt.Printf("langs[0]: %#v\n", n)

	//to change value of certain index inslice
	langs[1] = "Pythonistas"
	fmt.Printf("langs: %#v\n", langs)

	//similar to array
	//the difference is that slices length is flexible while array size is fixed and its zero value is nil
	//but how do we know the length of the slice?
	l := len(langs)
	fmt.Println("length of langs:", l)

	//to add more values into the slice
	langs = append(langs, "Fa", "Em")
	fmt.Printf("langs: %#v\n", langs)

	//Slicing , we can get certain value to certain value from the slice
	a := langs[0:2] // get value in slice from index 0-1
	fmt.Printf("a: %#v\n", a)

	//Get all values from index 0-last
	b := langs[0:len(langs)]
	fmt.Printf("n: %#v\n", b)

	//get values from index 0 to index 2 (3 values)
	r := langs[0:3]
	s := langs[:3]

	//get all values
	t := langs[0:]
	u := langs[:]

	fmt.Printf("r: %#v\n", r)
	fmt.Printf("s: %#v\n", s)
	fmt.Printf("t: %#v\n", t)
	fmt.Printf("u: %#v\n", u)

	//using slices as a prarameter for another function
	printSlice(langs)

	//change array to slice
	cords := [4]string{"A", "B", "C", "D"}
	printSlice(cords[:])

	//Zero value of slice
	// is nill when there is no underlying array
	var zero_langs []string
	fmt.Printf("zero langs: %#v\n", zero_langs)

	//is not nill when there is underlying array even when there's no value
	zero_langs_with_underlying_array := []string{}
	if zero_langs_with_underlying_array == nil {
		fmt.Println("Slice is nil")
	} else {
		fmt.Printf("Slice is NOT nil, value: %#v\n", zero_langs_with_underlying_array)
	}

}

func printSlice(ns []string) {
	fmt.Printf("printSlice: ns: %#v\n", ns)
}
