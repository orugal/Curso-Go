package main
import "fmt"
type User struct{
	edad int
	nombre,apellido,foto string
	documento string
}
/*Metodo para poner dentro de una función*/
func (this User) nombre_completo() string{
	return this.nombre + " " + this.apellido
}

///el * se pone para poder afectar la estructura a la cual se hace referencia, si no se le pasa el * crea una copia de la estructura
func (this *User) setName(n string){
	this.nombre = n
}

func main() {
	var datos User;

	datos.nombre = "Farez"
	datos.apellido = "Prieto"

	datos.setName("Sandra")
	fmt.Println(datos.nombre_completo())

}