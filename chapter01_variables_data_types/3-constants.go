package ch1

import "fmt"

// constantsDemo demonstrates the use of constants in go
func ConstantsDemo() {
	fmt.Println("=== constants ===")

	// simple constant
	const pi = 3.14159

	// typed constant
	const greeting string = "Hello, Go!"

	// multiple constants in a block
	const (
		minAge = 18
		maxAge = 60
	)

	// Constant expression
	const doublePi = pi * 2

	fmt.Println("Pi:", pi)
	fmt.Println("Greeting:", greeting)
	fmt.Println("Min Age:", minAge)
	fmt.Println("Max Age:", maxAge)
	fmt.Println("Double Pi:", doublePi)
}
