package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(v []int) {

	if len(v) > 1 {

		iPivot := rand.Intn(len(v))
		Pivot := v[iPivot]

		v[iPivot], v[len(v)-1] = v[len(v)-1], v[iPivot]

		i := 0
		for j := 0; j < len(v)-1; j++ {
			if v[j] <= Pivot {
				v[i], v[j] = v[j], v[i]
				i++
			}
		}

		v[i], v[len(v)-1] = v[len(v)-1], v[i]

		QuickSort(v[:i])
		QuickSort(v[i+1:])

	} else {
		return
	}
}

// ------------------ TESTE ----------------------
func main() {
	numeros := []int{21, 1, 13, 0, 233, 5, 3, 89, 8, 144, 2, 55, 1, 34, 377}
	fmt.Println("Antes:", numeros)

	QuickSort(numeros)

	fmt.Println("Depois:", numeros)
}
