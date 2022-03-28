package inputoutput

import "fmt"

func Demo() {
	var namn string
	fmt.Print("Ange namn:")
	fmt.Scanln(&namn)

	var age int
	fmt.Print("Ange ålder:")
	fmt.Scanln(&age)

	fmt.Printf("Hej %s du är %d år\n", namn, age)
}
