package main
import "fmt"

type Human struct{
	name string
}

type Tutor struct{
	Human
}


func main() {
	 tutor := Tutor{Human{"Farez"}}
	 fmt.Println(tutor.name)
}
