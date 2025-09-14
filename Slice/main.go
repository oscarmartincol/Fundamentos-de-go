package main

import (
	"fmt"
)

func main() {

	var intSlice []int

	intSlice = append(intSlice, 25, 14, 36)

	for i, v := range intSlice {
		fmt.Println(i, " value: ", v)
	}

	intSlice = append(intSlice, 52, 10, 6)

	for i, v := range intSlice {
		fmt.Println(i, " value: ", v)
	}

	intArray := [3]int{2, 7, 4}
	intSlice = intArray[1:2]

	for i, v := range intSlice {
		fmt.Println(i, " value: ", v)
	}

	intSlice = append(intSlice, intArray[0])

	for i, v := range intSlice {
		fmt.Println(i, " value: ", v)
	}
}
