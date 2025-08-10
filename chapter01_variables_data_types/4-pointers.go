package ch1

import "fmt"

// pointersDemo demonstrates how to use pointers in go
func PointersDemo() {
	fmt.Println("=== Pointers ===")

	// normal variable
	number := 10
	fmt.Println("Initial value:", number)

	// pointer to the variable (stores the memory address)
	var ptr *int = &number
	fmt.Println("Pointer address:", ptr)
	fmt.Println("Value via pointer:", *ptr)

	// modify the value via the pointer
	*ptr = 20
	fmt.Println("Modified value:", number)

	// Using 'new' to create a pointer to a zero value
	anotherPtr := new(int)
	*anotherPtr = 50
	fmt.Println("Value via another pointer:", *anotherPtr)

	// pointer to pointer
	var ptrToPtr **int = &ptr
	fmt.Println("Pointer to pointer address:", ptrToPtr)
	fmt.Println("Value via pointer to pointer:", **ptrToPtr)
}
