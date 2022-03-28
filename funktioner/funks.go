package funktioner

import (
	"fmt"
)

func add(x int, y int) (total int, err error) {
	err = nil
	total = x + y
	//return
	return total, nil
}

func withdraw(currentBalance int, belopp int) (newBalance int, errorText string) {
	errorText = ""
	newBalance = currentBalance
	if belopp > currentBalance {
		errorText = "Så mycket pengar har du inte"
	} else if belopp > 3000 {
		errorText = "Maxbelopp 3000 per uttag"
	} else {
		newBalance = currentBalance - belopp
	}
	return
}

func balance() {

	balance := 1000
	belopp := 1300
	errorText := ""
	balance, errorText = withdraw(balance, belopp)
	if len(errorText) > 0 {
		fmt.Println(errorText)
	}
}

func Run() {

	balance()

	// s, _ := add(12, 123)
	// fmt.Println(s)

	s, err := add(12, 123)
	if err == nil {
		fmt.Println(s)
	}

}
