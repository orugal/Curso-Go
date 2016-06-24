package main

import "fmt";

func main() {
	arreglo := [3] string{"Leshlee","Gabriela","Noelia"}
	for a:=0;a<len(arreglo);a++{
		fmt.Println(arreglo[a])
	}

	//multidimensionales
	var matriz [8][8] int;

	fmt.Print(matriz[0][2]);
}