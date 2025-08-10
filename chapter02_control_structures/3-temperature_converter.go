package ch2

import "fmt"

// temperatureConverterDemo converts temperatures between Celsius and Fahrenheit
func TemperatureConverterDemo() {
	var choice int
	fmt.Println("temperature converter")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Println("enter choice (1 or 2): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var celsius float64
		fmt.Println("enter temperature in celsius: ")
		fmt.Scan(&celsius)
		fahrenheit := (celsius * 9 / 5) + 32
		fmt.Printf("%.2fC = %.2fF\n", celsius, fahrenheit)
	case 2:
		var fahrenheit float64
		fmt.Println("enter temperature in fahrenheit: ")
		fmt.Scan(&fahrenheit)
		celsius := (fahrenheit - 32) * 5 / 9
		fmt.Printf("%.2fF = %.2fC\n", fahrenheit, celsius)
	default:
		fmt.Println("Invalid choice")
	}
}
