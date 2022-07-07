package main

import "fmt"

func main() {
	carro := []string{"pastanagues", "pebrots", "aigua", "formatge"}
	//Mira de recorrer i mostrar les dades através d'un for amb range.

	for i, v := range carro { //RANGE devuelve el índice y el valor de cada elemento
		fmt.Println(i, v)
	}

}
