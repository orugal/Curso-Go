package main

import "fmt"

func main() {
	/*
	* 1. Los punteros don direcciones de memoria.
	* 2. Las variables estan en una dirección de memória, los punteros nos dan la dirección y no el valor de la variable
	*/

	var x,y *int;
	entero :=5;

	x = &entero
	y = &entero

	*x = 6

	fmt.Println(x);
	fmt.Println(y);
	fmt.Println(entero);
}