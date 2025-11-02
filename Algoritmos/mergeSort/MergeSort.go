package main

import "fmt"

func Merge(v []int, e []int, d []int) {

	Ec, Dc := 0, 0

	for i := 0; i < len(e)+len(d); i++ {

		if Ec < len(e) && Dc < len(d) {
			if e[Ec] <= d[Dc] {
				v[i] = e[Ec]
				Ec++
			} else {
				v[i] = d[Dc]
				Dc++
			}

		} else {

			if Ec == len(e) {
				v[i] = d[Dc]
				Dc += 1
			} else {
				v[i] = e[Ec]
				Ec += 1
			}

		}
	}
}

func MergeSort(v []int) {

	if len(v) > 1 {

		mid := len(v) / 2
		left := make([]int, mid)
		right := make([]int, len(v)-mid)

		idV := 0

		for idE := range left {
			left[idE] = v[idV]
			idV++
		}
		for idD := range right {
			right[idD] = v[idV]
			idV++
		}

		MergeSort(left)
		MergeSort(right)
		Merge(v, left, right)

	} else {
		return
	}
}

// ------------------ TESTE ----------------------
func main() {
	numeros := []int{21, 1, 13, 0, 233, 5, 3, 89, 8, 144, 2, 55, 1, 34, 377}
	fmt.Println("Antes:", numeros)
	fmt.Println(len(numeros))

	MergeSort(numeros)

	fmt.Println("Depois:", numeros)
	fmt.Println(len(numeros))
}
