package main

import (
	"fmt"
)

type IQueue interface {
	Enqueue(value int)
	Dequeue() (int, error)
	Front() (int, error)
	IsEmpty() bool
	Size() int
}

type arrayQueue struct {
	data     []int
	length   int
	capacity int
}

func (aQ *arrayQueue) Init(size int) {
	aQ.data = make([]int, size)
	aQ.capacity = size
}

func (aQ *arrayQueue) DoubleCapacity() {
	var newAL = make([]int, 2*aQ.capacity)
	for i := 0; i < aQ.length; i++ {
		newAL[i] = aQ.data[i]
	}
	aQ.data = newAL
	aQ.capacity = 2 * aQ.capacity
}

func (aQ *arrayQueue) Size() int {
	return aQ.length
}

func (aQ *arrayQueue) IsEmpty() bool {
	if aQ.length == 0 {
		return true
	} else {
		return false
	}
}

func (aQ *arrayQueue) Front() (int, error) {
	if aQ.length != 0 {
		return aQ.data[0], nil
	} else {
		return -1, fmt.Errorf("A fila está vazia.")
	}
}

func (aQ *arrayQueue) Dequeue() (int, error) {
	if aQ.length == 0 {
		return -1, fmt.Errorf("A fila está vazia.")
	} else {
		var a = aQ.data[0]
		for i := 0; i < aQ.length; i++ {
			aQ.data[i] = aQ.data[i+1]
		}
		aQ.length--
		return a, nil
	}
}

func (aQ *arrayQueue) Enqueue(e int) {
	if aQ.length == aQ.capacity {
		aQ.DoubleCapacity()
	}
	aQ.data[aQ.length] = e
	aQ.length++
}

// ------------------ TESTE ----------------------
func main() {
	fmt.Println("=== Testando arrayQueue ===")

	q := &arrayQueue{}
	q.Init(3)

	// Teste Enqueue
	fmt.Println("\nAdicionando elementos 10, 20, 30")
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Println("Fila:", q.data[:q.length])
	fmt.Println("Tamanho:", q.Size())
	fmt.Println("Está vazia?", q.IsEmpty())

	// Teste Front
	f, err := q.Front()
	if err == nil {
		fmt.Println("Elemento da frente:", f)
	} else {
		fmt.Println("Erro:", err)
	}

	// Teste Enqueue além da capacidade (forçando expansão)
	fmt.Println("\nAdicionando 40 (forçando expansão)")
	q.Enqueue(40)
	fmt.Println("Fila:", q.data[:q.length])
	fmt.Println("Tamanho:", q.Size())
	fmt.Println("Capacidade:", q.capacity)

	// Teste Dequeue
	fmt.Println("\nRemovendo elementos da fila")
	for !q.IsEmpty() {
		val, err := q.Dequeue()
		if err == nil {
			fmt.Println("Dequeued:", val, "Fila restante:", q.data[:q.length])
		} else {
			fmt.Println("Erro:", err)
		}
	}

	// Teste Dequeue em fila vazia
	fmt.Println("\nTentando remover de fila vazia")
	_, err = q.Dequeue()
	if err != nil {
		fmt.Println("Erro:", err)
	}

	// Teste Front em fila vazia
	_, err = q.Front()
	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Println("\nFim dos testes!")
}
