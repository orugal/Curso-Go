package main

import "fmt";

/*MAKE*/

func main() {
	slice := make([]int,0,0)
	slice = append(slice,2)
	fmt.Println(cap(slice))
}