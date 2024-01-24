package lib

import "fmt"

// keyword 'var' ชื่อตัวแปร ประเภทตัวแปร = ค่าของตัวแปร

/*
 1. Variable scopes

declare variable at the top indicates package scope
every function inside this package will be able to see this variable
*/
var name_packageScope string = "Kidcat เองจ้า!!!"

func Variable() {
	/* delcare variable inside this function indicates function scope
	   this variable can be seen by only within this function */
	var name_functionScope string = "Kidcat เองจ้า!!!"
	fmt.Printf("My name is %s\n", name_functionScope)
	fmt.Printf("type of the variable is %T\n", name_functionScope)
}

func Variable2() {
	fmt.Printf("My name is %s\n", name_packageScope)
}

// 2. Ways to declare variable
func DifferentWaysToDeclareVariable() {
	/* There are 3 ways to declare variable
	   first two ways is the same and can be interchanged */
	var age int = 22
	// var age = 22

	/*  This 3rd way must be used inside a block scope (function) only
	    & must be a new variable name only */
	age2 := 22
	fmt.Printf("My age is %d and I will be %d this year", age, age2)
}

// default value for undeclared variable
var undeclared string

/*
	The zero value is:
	0 for numeric types
	false for boolean type
	"" (the emoty string) for strings.
*/

func UndeclaredVariable() {
	fmt.Printf("This variable default value is %q\n", undeclared)
}

/* Basic types in go
bool

string

int int8 int16 int32 int64
uint uint8 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
	// represents a Unicide code point
	// one character

float32 float64

complex64 complex128
\
*/

func BasicTypes() {
	var s string = "Meow"
	var i int = 2
	var f float64 = 3.4
	var b bool = true
	var r rune = 'c'

	fmt.Printf("string: %s\n", s)
	fmt.Printf("int: %v\n", i)
	fmt.Printf("float64: %v\n", f)
	fmt.Printf("bool: %v\n", b)
	fmt.Printf("rune: %v\n", r)

}
