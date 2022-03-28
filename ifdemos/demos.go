package ifdemos

import "fmt"

func SwitchDemo() {
	var age = 49
	namn := "Stefan"

	switch {
	case namn == "Stefan":
		fmt.Println("Ungdom")
	case age > 48:
		fmt.Println("Gamling")
	default:
		fmt.Println("Whatever")
	}

	var dayOfWeek = 6
	switch dayOfWeek {
	case 1, 2, 3:
		fmt.Println("Roliga dagar")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		{
			fmt.Println("Saturday")
			fmt.Println("Weekend. Yaay!")
		}
	case 7:
		{
			fmt.Println("Sunday")
			fmt.Println("Weekend. Yaay!")
		}
	default:
		fmt.Println("Invalid day")
	}

}

func Demos() {
	var age = 49

	namn := "Stefan"

	if age >= 18 {
		if namn == "Stefan" {
			fmt.Println("Du är inte vuxen")
		} else {
			fmt.Println("Du är vuxen")
		}
	} else if age >= 7 {
		fmt.Println("Du går i skolan")
	} else {
		fmt.Println("Du är en barnunge")
	}

}
