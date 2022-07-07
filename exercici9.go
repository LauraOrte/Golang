package main

/*Hem de realitzar una estructura condicional composta que contempli una nota y imprimeixi el següent:
- Menor de 5 "Suspes"
- Cinc "aprobat"
- Entre 6 i 8 (inclós) "Notable"
- En qualsevol altre cas "Excelent"
*/

import (
	"fmt"
)

func main() {
	num := 6

	if num < 5 {
		fmt.Println("Suspendido")
	} else if num == 5 {
		fmt.Println("Aprobado")
	} else if num < 9 {
		fmt.Println("Notable")
	} else {
		fmt.Println("Excelente")
	}
}
