package main

import (
	"fmt"
)

type any int //Para que el int ahora se llame any, se utiliza la palabra TYPE.

var anyActual any

func main() {
	anyActual = 2021
	fmt.Println(anyActual)      //2021
	fmt.Printf("%T", anyActual) //main.any, Para preguntarle el tipo se utiliza el PRINTF.
}
