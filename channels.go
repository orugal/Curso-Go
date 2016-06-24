package main

import (
"fmt"
//"time"
//"strings"
)

func main() {
	channel := make(chan string)

	go func(channel chan string){
		for{
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}(channel)

	for{
		msg := <- channel
		fmt.Println(msg)
	}
}