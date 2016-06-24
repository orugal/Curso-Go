package main
import "fmt"

type User interface {
	Permisos() int // 1 - 5
	Nombre() string
}

type Admin struct {
	nombre string
}

func (this Admin) Permisos() int {
	return 3
}

func (this Admin) Nombre() string {
	return this.nombre
}


func auth(user User) string{
	if user.Permisos() > 4 {
		return user.Nombre() + " tiene permisos de administrador"
	}else {
		return user.Nombre() + " no tiene permisos de administrador"
	}	
	//return ""
}

func main() {
	admin := Admin{"Farez"}
	fmt.Println(auth(admin))
}
