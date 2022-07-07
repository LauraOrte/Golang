package main

//Amb aquest exercici tens que realitzar una estructura condicional que representi un porter de discoteca que nomès deixa passar si ets major d'edat.
import (
	"fmt"
)

func main() {
	edad := 17

	// los CONDIONALES NO llevan parentesis para la condicion
	if edad >= 18 {
		fmt.Println("ok puedes pasar")
	} else {
		fmt.Println("Pa tu casa!")
	}
}
