package main

import "fmt"

func bubbleSort(v []int) {

	if len(v) > 1 {
		for i := 0; i < len(v)-1; i++ {
			for j := 0; j < len(v)-i-1; j++ {

				if v[j] > v[j+1] {
					aux := v[j]
					v[j] = v[j+1]
					v[j+1] = aux
					aux = 0
				}
			}
		}
	} else {
		fmt.Println("O vetor tem tamanho 1")
	}
}

// ------------------ TESTE ----------------------
func main() {
	numeros := []int{5, 3, 8, 1, 2, 7, 12, 95, 234, 39, 30192, 934, 124, 4, 3}
	fmt.Println("Antes:", numeros)

	bubbleSort(numeros)

	fmt.Println("Depois:", numeros)

}
