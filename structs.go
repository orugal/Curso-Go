package main

import "fmt"

type User struct{
	edad int
	nombre,apellido,foto string
	documento string
}


func main() {
	var datos User;
	datos2 := User{nombre:"Farez",apellido:"Prieto",edad:29,documento:"1030534849"}

	datos3 := new(User)	;
	fmt.Println(datos);
	fmt.Println(datos2);
	fmt.Println(datos3);
}