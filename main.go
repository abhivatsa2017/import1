package main

import (
	"fmt"
	"main/arrayer"
	"main/filedemo"
	"main/ifdemos"
	"main/inputoutput"
	"main/loopdemos"
	"main/oopdemo"
	"main/stackheap"
	"main/structs"
)

func oopFunc() {
	stefan := oopdemo.NewMonthly("Stefan", 12, "Test123", 12211, "Staden", 3000)
	oliver := oopdemo.NewHourly("Oliver", 12, "Test123", 12211, "Staden", 40)

	employees := []oopdemo.IPayable{}
	employees = append(employees, stefan)
	employees = append(employees, oliver)
	for _, emp := range employees {
		emp.CalculateSalary()
	}

	fmt.Printf("stefan: %v\n", stefan)
	fmt.Printf("oliver: %v\n", oliver)
}

func main() {
	filedemo.Demo()

	structs.Demo()
	oopFunc()
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
