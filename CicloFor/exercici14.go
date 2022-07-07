package main

import "fmt"

//parells del 0 al 20
func main() {

	alumnes := []string{"Anna", "Pep", "Iria", "Oscar"}
	//Realitza una estructora for que permeti recorrer tot l'array i finalitzi quan ja no hi hagin més elements per recorrer.

	for i := 0; i < len(alumnes); i++ {
		fmt.Println(alumnes[i])
	}
}
