package main

import (
	"fmt"
)

func main() {

	mapValues := make(map[int]string)
	mapValues[1] = "A"
	mapValues[10] = "B"
	mapValues[31] = "C"

	v, ok := mapValues[10] // Verificar si existe la posición del mapa.

	fmt.Println(v, ok)

	delete(mapValues, 10)

	for k, v := range mapValues {
		fmt.Println(k, v)
	}
}
