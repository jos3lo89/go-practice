package ch3

import (
	"fmt"
)

// SliceDemo demuestra el uso de slices
// los slices son mas flexibles y dinamicas que los array
func SliceDemo() {
	fmt.Println("\n=== Slices ===")

	// para crear un slice se usa la funcion make
	// creamos un slice de strings con longitud 3
	s := make([]string, 3)
	fmt.Println("slice vacio:", s)

	// asignar y obtener valores como en un array
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("slice asignado: ", s)
	fmt.Println("longitud: ", len(s))

	// <append> permite agregar nuevos elementos - <append> devuelve un nuevo slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("slice con append", s)
}
