package main

import (
	"fmt"
)

func main() {
	myInit := 4
	myIntPointer := &myInit
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)
	myFloat := 98.6
	myFloatPointer := &myFloat
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)
	myBool := true
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
	fmt.Println(*myBoolPointer)

	// actualizar un a localidad de memoria
	// se declara la variable y se asigna un valor
	myInit2 := 4
	// se imprime la variable
	fmt.Println(myInit2)
	//se declara la localidad de memoria de la variable myInit2
	myIntPointer2 := &myInit2
	fmt.Println(myIntPointer2)
	// se le asigna valor a esa localidad de memoria con el 8
	*myIntPointer2 = 8
	//imprime el valor de l localidad de memoria con "*"
	fmt.Println(myIntPointer2)
	//imprime el valor de la  variable directamente
	fmt.Println(myInit2)
	//Las  pruebas unitarias demostraron que una vez declararo el puntero y se actualiza el valor.... no se cambio la localidad de memoria
	//optimizando asi la memoria y no declrar valores en memoria por variable utilizada
}
