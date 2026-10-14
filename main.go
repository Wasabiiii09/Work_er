package main

import (
	"fmt"
)

func main(){
	fmt.Println("Welcome")
	number := 1
	for number < 11 {
		fmt.Println(number)
		number++ 
	}

	Texts := []string{"Kerim", "Damir"}
	fmt.Println(Texts[1])

	for name := range Texts {
		fmt.Println(name)
	}

}