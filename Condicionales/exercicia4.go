package main

import (
	"fmt"
)

var ciutat string = "Barcelona"

func main() {
	frase := fmt.Sprint("La ciutat de ", ciutat, " està a Catalunya") //SPRINT -> para unir elementos en un string y ponerlo en una variable
	fmt.Println(frase)                                                //La ciutat de Barcelona està a Catalunya

}
