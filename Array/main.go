package main

import (
	"fmt"
)

func main() {

	var intArray [5]int

	for i, v := range intArray {
		fmt.Println(i, " value: ", v)
	}
}
