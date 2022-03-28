package loopdemos

import "fmt"

func LoopDemos() {
	sum := 0

	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Printf("%d\n", sum)
	//fmt.Println()

	sum = 0
	counter := 1
	for counter <= 5 {
		sum += counter
		counter++
	}

	counter = 0
	for {
		counter++
		fmt.Printf("Nu är counter %d\n", counter)
		if counter > 10 {
			break
		}
	}

	// FOR EACH = använd RANGE
	stringarna := []string{"hello", "world"}
	for i := 0; i < len(stringarna); i++ {
		fmt.Println(stringarna[i])
	}
	for _, stringen := range stringarna {
		fmt.Printf("%s\n", stringen)
		//fmt.Printf("På position %d finns %s", index,stringen)
	}

}
