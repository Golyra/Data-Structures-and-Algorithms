package main

import (
	"fmt"
)

func CountingSort(v []int) {

	vmax := 0
	for i := range v {
		if v[i] > vmax {
			vmax = v[i]
		}
	}

	arr := make([]int, vmax+1)

	for j := range v {
		val := v[j]
		arr[val] += 1
	}

	pos := 0

	for g := range arr { // haha pirate arrrgh
		howMany := arr[g]
		if howMany > 0 {
			for range howMany {
				v[pos] = g
				pos++
			}
		}
	}
}

// ------------------ TESTE ----------------------
func main() {
	numeros := []int{21, 1, 13, 0, 233, 5, 3, 89, 8, 144, 2, 55, 1, 34, 377}
	fmt.Println("Antes:", numeros)

	CountingSort(numeros)

	fmt.Println("Depois:", numeros)
}
