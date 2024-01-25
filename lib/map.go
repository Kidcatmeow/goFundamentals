package main

import "fmt"

func main() {

	//map contain key and value
	//initialize map with keys: value
	status := map[int]string{
		200: "OK",
		300: "Permanent Redirect",
		400: "Bad Request",
		500: "Internal Server Error",
	}
	fmt.Printf("Print map: %#v\n", status)

	//count elements within the map
	l := len(status)
	fmt.Printf("length of the map: %d\n\n", l)

	//change value of certain key
	//when you enter new value nside the map , if the key already existed, the value will be replaced
	status[200] = "Okie"
	//if the key doesn't exist, new key and value will be added
	status[285] = "Pmsing"
	fmt.Printf("Print map: %#v\n", status)
	fmt.Printf("length of the map: %d\n\n", len(status))

	//how to read value from the map => assign value to a variable like arrays/slices
	value := status[200]
	fmt.Printf("%#v\n", value)

	//how to delete value from the map
	delete(status, 285)
	fmt.Printf("Print map: %#v\n", status)
	fmt.Printf("length of the map: %d\n\n", len(status))

	//in maps, an element with no value will always prints out as 0 in any given index
	v := status[66]
	fmt.Printf("value: %q \n\n", v)
	//so to check if the index really exist  , you can assign 2 values to the map (key and boolean)
	w, ok := status[66]
	if ok {
		fmt.Printf("value: %q \n\n", w)
	} else {
		fmt.Println("not found")
	}
	//could simplify to
	if w, ok := status[66]; ok {
		fmt.Printf("value: %q \n\n", w)
	} else {
		fmt.Println("not found")
	}

	//zero value of map is  (doesn't have space in memory yet)
	var m map[string]string
	//unless you add in {} , then the space of the map iwll be allocated within the memory
	var n map[string]string = make(map[string]string)
	//same as
	// var n = make(map[string]string)
	if n == nil {
		fmt.Printf("n is nil. value : %#v\n", n)
	} else {
		fmt.Printf("n is not nil. value: %#v\n", n)
	}

}
