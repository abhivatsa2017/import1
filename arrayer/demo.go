package arrayer

import "fmt"

type Person struct {
	name string
	age  int
}

func Demo() {

	//var personSlice = []Person{}
	var personSlice = make([]Person, 0, 5)
	personSlice = append(personSlice, Person{name: "Kerstin", age: 48})
	personSlice = append(personSlice, Person{name: "Stefan", age: 49})
	for _, person := range personSlice {
		fmt.Printf("%s är %d år\n", person.name, person.age)
	}

	var personerna = [5]Person{}
	p1 := Person{name: "Stefan", age: 49}
	personerna[0] = p1
	personerna[1] = Person{name: "Kerstin", age: 48}
	personerna[2] = Person{name: "Oliver", age: 13}
	personerna[3] = Person{name: "Fanny", age: 22}
	personerna[4] = Person{name: "Josefine", age: 20}

	for _, person := range personerna {
		fmt.Printf("%s är %d år\n", person.name, person.age)
	}

	var players = [...]string{"Stefan", "Mats", "Lisa"}
	for _, vardet := range players {
		fmt.Println(vardet)
	}

	//var arr1 = [...]int{1, 2, 3}
	var arr1 = [3]int{1, 2, 3}

	arr2 := [5]int{}

	arr2[0] = 123
	arr2[3] = 121

	fmt.Printf("%d\n%v\n", len(arr1), arr1)

}
