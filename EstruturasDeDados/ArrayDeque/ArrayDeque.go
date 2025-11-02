package main

import (
	"fmt"
)

type IDeque interface {
	EnqueueFront(value int) error
	EnqueueRear(value int) error
	DequeueFront() (int, error)
	DequeueRear() (int, error)
	Front() (int, error)
	Rear() (int, error)
	IsEmpty() bool
	Size() int
}

type arrayDeque struct {
	data     []int
	rear     int
	front    int
	capacity int
}

func (aD *arrayDeque) Init(capacity int) {
	aD.data = make([]int, capacity)
	aD.capacity = capacity
	aD.front = aD.capacity / 2
	aD.rear = aD.capacity / 2
}

func (aD *arrayDeque) EnqueueFront(e int) error {
	if (aD.front-1+aD.capacity)%aD.capacity != aD.rear {
		aD.data[aD.front] = e
		aD.front = (aD.front - 1 + aD.capacity) % aD.capacity

		return nil

	} else {
		return fmt.Errorf("A fila está cheia!")
	}
}

func (aD *arrayDeque) EnqueueRear(e int) error {

	if (aD.rear+1)%aD.capacity != aD.front {
		aD.rear = (aD.rear + 1) % aD.capacity
		aD.data[aD.rear] = e

		return nil

	} else {
		return fmt.Errorf("A fila está cheia!")
	}
}

func (aD *arrayDeque) DequeueFront() (int, error) {
	if aD.IsEmpty() {
		return -1, fmt.Errorf("A fila está vazia.")

	} else {
		aD.front = (aD.front + 1) % aD.capacity
		a := aD.data[aD.front]
		aD.data[aD.front] = 0

		return a, nil
	}
}

func (aD *arrayDeque) DequeueRear() (int, error) {
	behind := (aD.rear - 1 + aD.capacity) % aD.capacity

	if aD.IsEmpty() {
		return -1, fmt.Errorf("A fila está vazia.")

	} else {
		a := aD.data[aD.rear]
		aD.data[aD.rear] = 0
		aD.rear = behind

		return a, nil
	}
}

func (aD *arrayDeque) Size() int {
	if aD.rear > aD.front {
		return aD.rear - aD.front

	} else {
		return (aD.rear - aD.front) + aD.capacity
	}
}

func (aD *arrayDeque) IsEmpty() bool {
	if aD.front == aD.rear {
		return true
	} else {
		return false
	}
}

func (aD *arrayDeque) Front() (int, error) {
	if aD.front != aD.rear {
		return aD.data[(aD.front+1)%aD.capacity], nil
	} else {
		return -1, fmt.Errorf("A fila está vazia.")
	}
}

func (aD *arrayDeque) Rear() (int, error) {
	if aD.front != aD.rear {
		return aD.data[aD.rear], nil
	} else {
		return -1, fmt.Errorf("A fila está vazia.")
	}
}

// ------------------ TESTE ----------------------
func main() {
	q := &arrayDeque{}
	q.Init(6) // capacidade 6

	fmt.Println(">>> Teste EnqueueRear")
	_ = q.EnqueueRear(10)
	_ = q.EnqueueRear(20)
	_ = q.EnqueueRear(30)
	fmt.Println("Fila:", q.data, "Size:", q.Size())

	fmt.Println("\n>>> Teste EnqueueFront")
	_ = q.EnqueueFront(5)
	_ = q.EnqueueFront(2)
	fmt.Println("Fila:", q.data, "Size:", q.Size())

	fmt.Println("\n>>> Verificando Front e Rear")
	f, _ := q.Front()
	r, _ := q.Rear()
	fmt.Println("Front:", f, "Rear:", r)

	fmt.Println("\n>>> Teste DequeueFront")
	val, _ := q.DequeueFront()
	fmt.Println("Saiu:", val, "Fila:", q.data, "Size:", q.Size())
	val, _ = q.DequeueFront()
	fmt.Println("Saiu:", val, "Fila:", q.data, "Size:", q.Size())

	fmt.Println("\n>>> Teste DequeueRear")
	val, _ = q.DequeueRear()
	fmt.Println("Saiu:", val, "Fila:", q.data, "Size:", q.Size())
	val, _ = q.DequeueRear()
	fmt.Println("Saiu:", val, "Fila:", q.data, "Size:", q.Size())

	fmt.Println("\n>>> Estado final")
	fmt.Println("Fila:", q.data, "Size:", q.Size(), "IsEmpty:", q.IsEmpty())

	fmt.Println("\n>>> Teste fila vazia")
	val, err := q.DequeueFront()
	fmt.Println("Saiu:", val, "Fila:", q.data, "Size:", q.Size())

	_, err = q.DequeueRear()
	if err != nil {
		fmt.Println("Erro esperado:", err)
	}
	_, err = q.DequeueFront()
	if err != nil {
		fmt.Println("Erro esperado:", err)
	}
}
