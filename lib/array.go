package lib

import "fmt"

// เป็นเหมือนห้องแถว
// index always start with 0
// อันนี้คือ array ห้องแถวที่ชื่อว่า langs ที่เป็นห้องแถว 4 ห้อง
var Langs = [4]string{}

// to insert values into the array simply
var LangsValues = [4]string{"golang", "java", "python"}

func ShowAll(ns [4]string) {
	fmt.Printf("show all : %#v\n", ns)
}

func array() {
	var num = [4]int{1, 2, 3, 4}
	remove_ele(num, 1)
}

func remove_ele(originalArray [4]int, index int) {

	// initialize the index of the element to delete
	i := index

	// check if the index is within array bounds
	if i < 0 || i >= len(originalArray) {
		fmt.Println("The given index is out of bounds.")
	} else {
		// delete an element from the array
		newLength := 0
		for index := range originalArray {
			if i != index {
				originalArray[newLength] = originalArray[index]
				newLength++
			}
		}

		// reslice the array to remove extra index
		newArray := originalArray[:newLength]
		fmt.Println("The new array is:", newArray)
	}
}
