package structs

import "fmt"

// Stor bokstav = synliga från andra packages

type Address struct {
	StreetAddress string
	PostalCode    int
	City          string
}

type Person struct {
	Name      string
	Age       int
	Addressen Address
}

type Animal struct {
	Name      string
	Addressen Address
}

func createPerson() *Person {
	p1 := Person{Name: "Stefan", Age: 49,
		Addressen: Address{StreetAddress: "Testg", PostalCode: 123, City: "Test"}}
	return &p1
}

func Demo() {
	p3 := createPerson()
	fmt.Println(p3)
	p1 := Person{Name: "Stefan", Age: 49,
		Addressen: Address{StreetAddress: "Testg", PostalCode: 123, City: "Test"}}

	p2 := Person{Name: "Stefan", Age: 49,
		Addressen: Address{StreetAddress: "Testg", PostalCode: 123, City: "Test"}}

	a1 := Animal{Name: "Musse",
		Addressen: Address{StreetAddress: "Testg", PostalCode: 123, City: "Test"}}

	fmt.Println(a1)
	if p1 == p2 {
		fmt.Println("P1 och P2 är lika")
	}

}
