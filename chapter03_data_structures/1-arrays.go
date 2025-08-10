// Package ch3
package ch3

import "fmt"

// ArraysDemo los arrays tienen un tamano fijo y no pueden cambiar de longitud una vez declarada
func ArraysDemo() {
	fmt.Println("=== Arrays ===")

	// declaracion  de un array  de 5 enteros
	// por defecto se inicializa e 0
	var a [5]int
	fmt.Println("array vacio: ", a)

	// asignar y obtener usando el indice
	a[4] = 33
	fmt.Println("asignado: ", a)
	fmt.Println("obetenido: ", a[4])

	// declarar e inicializar un array
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declarado e inicializado:", b)

	// inferiendo la longitud de un array
	c := [...]int{1, 2, 3}
	fmt.Println("longitud inferida:", c, "longitud: ", len(c))

	// array multi-dimencionales
	var twoD [2][3]int

	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("array 2d", twoD)
}
