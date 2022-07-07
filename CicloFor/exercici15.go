package main

import "fmt"

//parells del 0 al 20
func main() {

	//Dibuixa un cuadrat amb un únic * i que tingui 10 unitats d'ample i 10 d'alt.
	simbol := "*"

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print(simbol)
		}
		fmt.Println("")
	}
}
