package main

import "fmt"

//parells del 0 al 20
func main() {

	i := 0

	for i <= 20 {
		if i%2 == 0 {
			fmt.Println(i)
		}
		i++

	}
}
