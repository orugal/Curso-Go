package main
import("fmt")

func main(){
	var nro int;
	fmt.Println("Ingresa un número:");
	fmt.Scanf("%d\n",&nro);
	pintaTabla(nro);
}
func pintaTabla(nro int){
	var a,b int;
	for a=0; a<=nro; a++ {
		for b=0;b<=nro;b++ {
			var c = (b % 2)
			if(c == 0){
				fmt.Print("  *  ");
			}else{
				fmt.Print("  o  ");	
			}
		}
		fmt.Println();
		fmt.Print("\n");
	}
}