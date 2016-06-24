package main

import "fmt"

func main() {

	array1 := []int{1,2,3,4}
	copia  := make([]int,len(array1)*2)

	copy(copia,array1);
	fmt.Println(array1)
	fmt.Println(copia)
}