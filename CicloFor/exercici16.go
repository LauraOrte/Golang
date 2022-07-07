package main

import "fmt"

//parells del 0 al 20
func main() {

	//Mira d'obtenir el total de la suma dels numeros parells de la col.lecció que indiquem i només emprant una estructura for sense atributs.
	nums := []int{3, 6, 12, 7, 5, 8, 2}

	i, suma := 0, 0

	for {
		i++
		if i == len(nums) {
			break
		}
		if nums[i]%2 == 0 {
			suma += nums[i]
		} else {
			continue //continue sirve para saltar una iteración
		}

	}

	fmt.Println(suma)
}
