package ch2

import "fmt"

// switchDemo demostrates the use of switch statement in go

func SwitchDemo() {
	day := "monday"

	switch day {
	case "monday":
		fmt.Println("start of the work week")
	case "friday":
		fmt.Println("almost week")
	case "saturday":
		fmt.Println("weekend")
	default:
		fmt.Println("midweek day")
	}

	// switch without an expresssion (acts like if-else)
	num := 15
	switch {
	case num < 0:
		fmt.Println("negative number")
	case num == 0:
		fmt.Println("zero")
	default:
		fmt.Println("positive number")
	}
}
