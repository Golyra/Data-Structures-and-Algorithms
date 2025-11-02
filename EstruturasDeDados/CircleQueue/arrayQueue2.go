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
	rear     int
	front    int
	capacity int
}

func (aQ *arrayQueue) Init(capacity int) {
	aQ.data = make([]int, capacity)
	aQ.capacity = capacity
	aQ.front = -1
	aQ.rear = -1
}

func (aQ *arrayQueue) Size() int {
	if aQ.rear > aQ.front {
		return aQ.rear - aQ.front

	} else {
		return (aQ.rear - aQ.front) + aQ.capacity
	}
}

func (aQ *arrayQueue) IsEmpty() bool {
	if aQ.front == aQ.rear {
		return true
	} else {
		return false
	}
}

func (aQ *arrayQueue) Front() (int, error) {
	if aQ.front != aQ.rear {
		return aQ.front + 1, nil
	} else {
		return -1, fmt.Errorf("A fila está vazia.")
	}
}

func (aQ *arrayQueue) Dequeue() (int, error) {
	if aQ.IsEmpty() {
		return -1, fmt.Errorf("A fila está vazia.")

	} else {
		aQ.front = (aQ.front + 1) % aQ.capacity
		a := aQ.data[aQ.front]
		aQ.data[aQ.front] = 0

		return a, nil
	}
}

func (aQ *arrayQueue) Enqueue(e int) error {

	if (aQ.rear+1)%aQ.capacity != aQ.front {

		if aQ.rear == -1 && aQ.front == -1 {
			aQ.rear = 0
			aQ.data[aQ.rear] = e

		} else {
			aQ.rear = (aQ.rear + 1) % aQ.capacity
			aQ.data[aQ.rear] = e
		}

		return nil
	} else {
		return fmt.Errorf("A fila está cheia!")
	}
}

// ------------------ TESTE ----------------------
func main() {
	q := &arrayQueue{}
	q.Init(5)

	fmt.Println(">>> Teste Enqueue")
	fmt.Println("Enfileirando 10, 20, 30")
	_ = q.Enqueue(10)
	_ = q.Enqueue(20)
	_ = q.Enqueue(30)
	fmt.Println("Fila:", q.data, "Size:", q.Size(), "Front:", q.front, "Rear:", q.rear)

	fmt.Println("\n>>> Teste Dequeue")
	val, _ := q.Dequeue()
	fmt.Println("Saiu:", val)
	val, _ = q.Dequeue()
	fmt.Println("Saiu:", val)
	fmt.Println("Fila:", q.data, "Size:", q.Size(), "Front:", q.front, "Rear:", q.rear)

	fmt.Println("\n>>> Teste Enqueue após Dequeue")
	_ = q.Enqueue(40)
	_ = q.Enqueue(50)
	_ = q.Enqueue(60)
	fmt.Println("Fila:", q.data, "Size:", q.Size(), "Front:", q.front, "Rear:", q.rear)

	fmt.Println("\n>>> Teste Fila Cheia")
	err := q.Enqueue(70)
	if err != nil {
		fmt.Println("Erro esperado:", err)
	}

	fmt.Println("\n>>> Teste Dequeue até esvaziar")
	for !q.IsEmpty() {
		val, _ := q.Dequeue()
		fmt.Println("Saiu:", val, "Fila:", q.data, "Front:", q.front, "Rear:", q.rear)
	}

	fmt.Println("\n>>> Teste Fila Vazia")
	_, err = q.Dequeue()
	if err != nil {
		fmt.Println("Erro esperado:", err)
	}
}
