package main

import "fmt"

//Heu de realitzar una estructura condicional Switch que contempli els seguents punts basant-se amb l'exemple del porter de discoteca:
// 1. Podra accedir si és major d'edat"
// 2. Si és major de 16 anys (inclòs) i és Dijous.
func main() {

	edat := 21
	dia := "Dijous"

	switch {
	case edat > 18, (edat > 16 && dia == "Dijous"):
		fmt.Println("Entra")
	default:
		fmt.Println("no puedes pasar")

	}

}
