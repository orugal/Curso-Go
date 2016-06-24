package main

import ("fmt"
		"strconv"
)


func main(){
	edad := 29
	otraEdad := "29"
	//strconv.Itoa  Convierte un int a String pero se debe incluir el paquete strconv
	edad_str := strconv.Itoa(edad)
	//strconv.Atoi  Convierte un int a String pero se debe incluir el paquete strconv
	edad_int,_ := strconv.Atoi(otraEdad);

	fmt.Println("Mi edad es: "+edad_str);
	fmt.Println(edad_int);
}