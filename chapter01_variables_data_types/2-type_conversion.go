package ch1

import (
	"fmt"
	"strconv"
)

// typeConversionDemo shows how to convert between types in go
func TypeConversionDemo() {
	fmt.Println("=== type conversion ===")

	// integer to float
	var i int = 42
	var f float64 = float64(i)
	fmt.Println("int to float:", f)

	// float to int (decimal part is truncated)
	var g float64 = 3.99
	var j int = int(g)
	fmt.Println("float to int:", j)

	// integer to string
	s := strconv.Itoa(i)
	fmt.Println("int to striung:", s)

	//string to integer (with error handling)
	numStr := "100"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("error converting string to int:", err)
	} else {
		fmt.Println("string to int:", num)
	}

	// string to float
	floatStr := "3.1415"
	floatNum, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("error converting string to float:", err)
	} else {
		fmt.Println("string to float:", floatNum)
	}
}
