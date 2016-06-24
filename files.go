package main

import(
	"fmt"
	"io/ioutil"
	"strings"
	"bufio"
	"os"
)

type Pala struct{
	palabra string
	cantLetras int
}


func main(){
	leeArchivos2()
}


//Forma Nro 1 de leer archivos
func leeArchivos(){
	datos,err := ioutil.ReadFile("./hola.txt")
	var datosFinal string
	var cantidadCar int
	//var array[] int
	if err != nil {
		fmt.Println("Hubo un error")
	}else{
		datosFinal = string(datos)
		pedazos := strings.Split(datosFinal," ")
		for _, palabras := range(pedazos){
			cantidadCar = len(palabras)
			//array = cantidadCar
			obj := Pala{palabra:palabras,cantLetras:cantidadCar}
			fmt.Println(obj)
		}
	}
}


//forma de leer archivo linea por linea
func leeArchivos2(){
	archivo,err := os.Open("./hola.txt")//abrir archivo
	if err != nil {
		fmt.Println("Hubo un error")
	}else{
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan(){
			linea := scanner.Text()
			fmt.Println(linea)
		}
	}
}