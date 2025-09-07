package main

import (
	"fmt"
)

type List interface {
	Size() int
	Get(index int) (int, error)
	Add(e int)
	AddOnIndex(e int, index int) error
	Remove(index int) error
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

func (aL *arrayList) Remove(index int) error {
	if index < 0 && index > aL.length {
		return fmt.Errorf("Index out of range: %d", index)
	}
	for i := index; i < aL.length; i++ {
		aL.data[i] = aL.data[i+1]
	}
	aL.length--
	return nil
}

// ------------------ TESTE ----------------------
func main() {
	fmt.Println("=== Testando arrayList ===")

	aL := &arrayList{}
	aL.Init(3)

	// Teste Add
	fmt.Println("\nAdicionando elementos...")
	aL.Add(10)
	aL.Add(20)
	aL.Add(30)
	fmt.Println("Lista:", aL.data[:aL.length], "Tamanho:", aL.Size(), "Capacidade:", aL.capacity)

	// Teste expansão de capacidade
	aL.Add(40)
	fmt.Println("Após expansão:", aL.data[:aL.length], "Tamanho:", aL.Size(), "Capacidade:", aL.capacity)

	// Teste Get
	val, err := aL.Get(2)
	if err == nil {
		fmt.Println("Elemento no índice 2:", val)
	} else {
		fmt.Println(err)
	}

	// Teste AddOnIndex
	fmt.Println("\nInserindo 15 no índice 1...")
	aL.AddOnIndex(15, 1)
	fmt.Println("Lista:", aL.data[:aL.length])

	// Teste Remove
	fmt.Println("\nRemovendo índice 2...")
	aL.Remove(2)
	fmt.Println("Lista:", aL.data[:aL.length])

	// Teste Pop
	fmt.Println("\nUsando Pop...")
	aL.Pop()
	fmt.Println("Lista:", aL.data[:aL.length])

	fmt.Println("\nFim dos testes!")
}
