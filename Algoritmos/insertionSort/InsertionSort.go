package main

import "fmt"

func insertionSort(v []int) {

	for i := 1; i < len(v); i++ {

		key := v[i]
		pos := i

		for j := i; j >= 0; j-- {
			if v[j] > key {
				pos = j
				v[j+1] = v[j]
			}
		}

		v[pos] = key
	}
}

// ------------------ TESTE ----------------------
func main() {
	numeros := []int{21, 1, 13, 0, 233, 5, 3, 89, 8, 144, 2, 55, 1, 34, 377}
	fmt.Println("Antes:", numeros)

	insertionSort(numeros)

	fmt.Println("Depois:", numeros)
}
