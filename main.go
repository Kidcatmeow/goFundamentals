package main

import (
	"KBTG_Reskill/lib"
	"fmt"
)

func main() {
	// lib.Variable()
	// lib.Different_print()
	// lib.UndeclaredVariable()
	// lib.BasicTypes()
	// lib.Const()
	// lib.If()
	// lib.Switch()
	lib.Info("Watsamon", "Yoyo", 22)
	day_from_Today := lib.Today()
	fmt.Println(day_from_Today)

	a, b := lib.Swap("one", "two")
	fmt.Println(a, b)

	c := lib.Plus(2, 5)
	fmt.Println(c)
	fmt.Printf("type: %T\n", lib.Plus)

	//We basically send Plus and Monus function into Cal
	lib.Cal(lib.Plus)
	lib.Cal(lib.Minus)
}
