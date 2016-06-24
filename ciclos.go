package main

import "fmt";

func main() {
	var	a int;
	for a=1;a<=10;a++{
		fmt.Printf("Contando --------- %v \n",a)
		if a==5{
			break;
		}
	}
}