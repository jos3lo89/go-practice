package ch2

import (
	"fmt"
)

// IfElseDemo demostrates the use of if/else conditional statements in go
func IfElseDemo() {
	number := 10

	// basic if statement
	if number > 0 {
		fmt.Println("the number is positive")
	}

	// if-eelse statement
	if number%2 == 0 {
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}

	// if-else if-else statement
	age := 20
	if age < 13 {
		fmt.Println("you are a child")
	} else if age < 18 {
		fmt.Println("you are a teenager")
	} else {
		fmt.Println("you are an adult")
	}

	// declare variable and if-else statement
	if isTrue := true; isTrue {
		fmt.Print("is true")
	} else {
		fmt.Println("is false")
	}
}
