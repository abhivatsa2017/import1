package main

import (
	"fmt"
	"main/arrayer"
	"main/ifdemos"
	"main/inputoutput"
	"main/loopdemos"
	"main/stackheap"
	"main/structs"
)

func main() {
	structs.Demo()
	stackheap.Demo()
	arrayer.Demo()
	for {
		var sel int
		_, err := fmt.Scanln(&sel)
		if err == nil {
			break
		}
		var clearbuf string
		fmt.Scanln(&clearbuf)
		fmt.Printf("Inte valid nummer")
	}

	ifdemos.SwitchDemo()
	loopdemos.LoopDemos()

	age := 12
	fmt.Printf("Hej %d du är %T år\n", age, age)

	inputoutput.Demo()

	// var age int
	// age = 12

	// var tal1 int8
	// var tal2 int16
	// var tal3 int64

	// tal1 = 1
	// tal2 = 12

	// tal3 = tal1 + tal2

	// namn := "stefan	holmberg"
	// namn = strings.ToTitle(namn)

	// // var pi float64
	// // pi = 3.1415

	// var salary = 105.50

	// fmt.Printf("Hejsan hoppsan %s du är %d %f år\n", namn, age, salary)
}
