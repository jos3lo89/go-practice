package ch1

import "fmt"

// variablesDemo demonstrates how to declare and use variables in go
func VariablesDemo() {
	fmt.Println("=== variables and data types ===")

	// short variable declaration (type is inferred)
	agency := "Lagarto inc"
	name := "Jos3lo33"

	// explicit type declaration
	var age int = 30

	// multiple variable declaration
	var height, weight float64 = 3.33, 9.99

	// zero values: variables without initialization get default values
	var city string   // default ""
	var score int     // default 0
	var isActive bool // default false

	fmt.Println("Agency:", agency)
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height, "m")
	fmt.Println("Weight:", weight, "kg")
	fmt.Println("City:", city)
	fmt.Println("Score:", score)
	fmt.Println("Active:", isActive)
}
