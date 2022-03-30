package structs

import (
	"github.com/cheynewallace/tabby"
)

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

var i = 12

func Demo() {

	// myDictionary := map[string]int{}
	// myDictionary2 := make(map[string]int)
	// myDictionary3 := make(map[string]int, 10)

	// fmt.Println(myDictionary)
	// fmt.Println(myDictionary2)
	// fmt.Println(myDictionary3)

	// p3 := createPerson()
	// fmt.Println(p3)
	p1 := Person{Name: "Stefan", Age: 49,
		Addressen: Address{StreetAddress: "Testg", PostalCode: 123, City: "Test"}}

	p2 := Person{Name: "Stefan", Age: 49,
		Addressen: Address{StreetAddress: "Testg", PostalCode: 123, City: "Test"}}

	t := tabby.New()
	t.AddHeader("Namn", "Ålder", "City")
	t.AddLine(p1.Name, p1.Age, p1.Addressen.City)
	t.AddLine(p2.Name, p2.Age, p2.Addressen.City)
	t.Print()
}
