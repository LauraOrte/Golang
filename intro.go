package main

import "fmt"

func main() {
	var a int
	/*var b string
	var c float64
	var d bool*/

	fmt.Printf("var a %T =%+v\n", a, a) // Dime el tipo (%T), y el valor(%+v) de la variable a para el tipo, y de la variable a para el valor. Salto de linia (\n)

	//LO MISMO PARA TODOS PA VERLO (b, c, d)
	//Verbos %T, %+v,

	//TABLA MULTIPLICAR;
	num1 := 2
	num2 := 3

	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)

}
