package ch2

import "fmt"

// forLoopDemo demostrates different types of for loops in go

func ForLoopDemo() {
	// classic for loop
	fmt.Println("classic for loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while-style loop
	fmt.Println("\nwhile-style loop")
	count := 0
	for count < 3 {
		fmt.Println(count)
		count++
	}

	// infinite loop (use break to stop)
	fmt.Println("\ninfinite loop with break")
	num := 1
	for {
		fmt.Println(num)
		num++
		if num > 3 {
			break
		}
	}

	// loop with continue
	fmt.Println("\nloop with continue")
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue // skip even number
		}
		fmt.Println(i)
	}
}
