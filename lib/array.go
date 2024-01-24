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
