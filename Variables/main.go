package main

// Librerias
import (
	"fmt" // imprimir pantalla
	"strconv"
	"unsafe"
)

func main() {

	// Declaración de variable
	var myVarInt int
	fmt.Println("My variable is: ", myVarInt)

	fmt.Println("Dirección de la variable: ", &myVarInt)

	{
		const myIntConst int = 12
		myScopeVariable := 40
		fmt.Println("Valor de la constante: ", myIntConst)
		fmt.Println("Valor del Scope: ", myScopeVariable)
	}

	/*Tamaño de las variables*/

	var my8BitsuintVar int = 12
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my8BitsuintVar, my8BitsuintVar, unsafe.Sizeof(my8BitsuintVar), unsafe.Sizeof(my8BitsuintVar)*8)
	MyIntStr := fmt.Sprintf("%d", my8BitsuintVar) // Convierte entero a string
	fmt.Printf("Type: %T, Value: %s \n", MyIntStr, MyIntStr)

	floatVar := 23.5
	floatStr := strconv.FormatFloat(floatVar, 'f', -1, 64) // Otra forma de convertir String.
	fmt.Printf("Type: %T, Value: %s \n", floatStr, floatStr)

	/*Convertir String a numero*/
	strIntVar, err := strconv.Atoi("15")
	fmt.Printf("Type: %T, Value: %d, err: %v \n", strIntVar, strIntVar, err)

	strIntVar2, err := strconv.ParseInt("17", 10, 64)
	fmt.Printf("Type: %T, Value: %d, err: %v \n", strIntVar2, strIntVar2, err)

	strIntVar3, err := strconv.ParseFloat("-11.2", 64)
	fmt.Printf("Type: %T, Value: %f, err: %v \n", strIntVar3, strIntVar3, err)

}
