package main

import "fmt"

func selectionSort(v []int) {

	lowest := v[0]

	for i := 0; i < len(v); i++ {

		lowest = v[i]
		pos := i

		for j := i; j < len(v); j++ {

			if v[j] < lowest {
				lowest = v[j]
				pos = j
			}
		}

		v[pos] = v[i]
		v[i] = lowest
	}
}

// ------------------ TESTE ----------------------
func main() {

	numeros := []int{5, 3, 8, 1, 2}
	fmt.Println("Antes:", numeros)

	selectionSort(numeros)

	fmt.Println("Depois:", numeros)

}
