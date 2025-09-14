package main

import (
	"fmt"
)

func main() {

	myIntvar := 16

	fmt.Println(myIntvar < 16 || myIntvar > 8)    // OR
	fmt.Println(myIntvar < 16 && myIntvar > 8)    // AND
	fmt.Println(!(myIntvar < 16 || myIntvar < 8)) // NOT
}
