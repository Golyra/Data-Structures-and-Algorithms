package main

import (
	"fmt"
)

type IList interface {
	Add(value int)
	AddOnIndex(value int, index int) error
	RemoveOnIndex(index int) error
	Get(index int) (int, error)
	Set(value int, index int) error
	Size() int
}

type arrayList struct {
	data     []int
	length   int
	capacity int
}

func (aL *arrayList) Init(size int) {
	aL.data = make([]int, size)
	aL.capacity = size
}

func (aL *arrayList) DoubleCapacity() {
	var newAL = make([]int, 2*aL.capacity)
	for i := 0; i < aL.length; i++ {
		newAL[i] = aL.data[i]
	}
	aL.data = newAL
	aL.capacity = 2 * aL.capacity
}

func (aL *arrayList) Size() int {
	return aL.length
}

func (aL *arrayList) Get(index int) (int, error) {
	if index >= 0 && index < aL.length {
		return aL.data[index], nil
	} else {
		return -1, fmt.Errorf("Error: index out of range (%d)", index)
	}
}

func (aL *arrayList) Add(e int) {
	if aL.length == aL.capacity {
		aL.DoubleCapacity()
	}
	aL.data[aL.length] = e
	aL.length++
}

func (aL *arrayList) AddOnIndex(e int, index int) error {
	if index < 0 && index > aL.length {
		return fmt.Errorf("Index out of range: %d", index)
	}
	if aL.length == aL.capacity {
		aL.DoubleCapacity()
	}
	for i := aL.length; i > index; i-- {
		aL.data[i] = aL.data[i-1]
	}
	aL.data[index] = e
	aL.length++
	return nil
}

func (aL *arrayList) Pop() {
	aL.data[aL.length] = 0
	aL.length--
}

func (aL *arrayList) RemoveOnIndex(index int) error {
	if index < 0 && index > aL.length {
		return fmt.Errorf("Index out of range: %d", index)
	}
	for i := index; i < aL.length; i++ {
		aL.data[i] = aL.data[i+1]
	}
	aL.length--
	return nil
}

func (aL *arrayList) Set(value int, index int) error {
	if index < 0 || index >= aL.length {
		return fmt.Errorf("índice fora do intervalo")
	}
	aL.data[index] = value
	return nil
}

func rev_bin_search(val int, list []int, start int, end int) int {
	if start > end {
		return -1 // não encontrado
	}
	mid := (start + end) / 2

	if list[mid] == val {
		return mid
	} else if list[mid] > val {
		return rev_bin_search(val, list, start, mid-1)
	} else {
		return rev_bin_search(val, list, mid+1, end)
	}
}

// ------------------ TESTE ----------------------
func main() {
	index := rev_bin_search(42, []int{10, 20, 30, 40, 42, 50}, 0, 5)
	fmt.Println("Índice encontrado:", index)
}
