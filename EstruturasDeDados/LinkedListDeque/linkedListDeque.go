package main

import (
	"fmt"
)

type linkedListDeque struct {
	head   *Node
	length int
	tail   *Node
}

type Node struct {
	prev *Node
	val  int
	next *Node
}

type IDeque interface {
	EnqueueFront(value int)
	EnqueueRear(value int)
	DequeueFront() (int, error)
	DequeueRear() (int, error)
	Front() (int, error)
	Rear() (int, error)
	IsEmpty() bool
	Size() int
}

func (LLD *linkedListDeque) EnqueueFront(value int) {
	newNode := &Node{val: value, next: LLD.head}
	if LLD.head == nil {
		LLD.head = newNode
		LLD.tail = newNode
	} else {
		LLD.head.prev = newNode
		newNode.next = LLD.head
		LLD.head = newNode
	}
	LLD.length++
}

func (LLD *linkedListDeque) EnqueueRear(value int) {
	newNode := &Node{val: value, prev: LLD.tail}
	if LLD.head == nil {
		LLD.head = newNode
		LLD.tail = newNode
	} else {
		LLD.tail.next = newNode
		newNode.prev = LLD.tail
		LLD.tail = newNode
	}
	LLD.length++
}

func (LLD *linkedListDeque) DequeueFront() (int, error) {
	var aux *Node

	if LLD.IsEmpty() {
		return -1, fmt.Errorf("A lista está vazia")
	} else {
		aux = LLD.head
		a := aux.val
		LLD.head = aux.next
		aux.val = 0
		aux.prev = nil
		aux.next = nil
		LLD.length--

		return a, nil
	}
}

func (LLD *linkedListDeque) DequeueRear() (int, error) {

	if LLD.IsEmpty() {
		return -1, fmt.Errorf("A lista está vazia")
	} else {
		var aux *Node
		var a int

		aux = LLD.tail
		a = aux.val
		LLD.tail = LLD.tail.prev
		aux.prev = nil
		aux.next = nil
		aux.val = 0
		LLD.length--

		return a, nil
	}
}

func (LLD *linkedListDeque) IsEmpty() bool {
	if LLD.length == 0 {
		return true
	} else {
		return false
	}
}

func (LLD *linkedListDeque) Size() int {
	return LLD.length
}

// Pra facilitar a visualização dos testes
func (LLD *linkedListDeque) ImprimirLista() {

	if (LLD.head) != nil {
		aux := (LLD.head)
		fmt.Printf("[")

		for aux != nil {
			fmt.Printf("%d", aux.val)
			aux = aux.next
			if aux != nil {
				fmt.Printf(", ")
			}
		}
		fmt.Println("]")
	} else {
		fmt.Println("A lista está vazia!")
	}
}

// =========================================== TESTES ===========================================
func main() {
	q := &linkedListDeque{}

	fmt.Println(">>> Teste inicial")
	fmt.Println("IsEmpty:", q.IsEmpty(), "Size:", q.Size())
	q.ImprimirLista()

	fmt.Println("\n>>> Teste EnqueueRear")
	q.EnqueueRear(10)
	q.EnqueueRear(20)
	q.EnqueueRear(30)
	q.ImprimirLista()
	fmt.Println("Size:", q.Size())

	fmt.Println("\n>>> Teste EnqueueFront")
	q.EnqueueFront(5)
	q.EnqueueFront(2)
	q.ImprimirLista()
	fmt.Println("Size:", q.Size())

	fmt.Println("\n>>> Teste DequeueFront")
	val, _ := q.DequeueFront()
	fmt.Println("Saiu:", val)
	q.ImprimirLista()
	fmt.Println("Size:", q.Size())

	val, _ = q.DequeueFront()
	fmt.Println("Saiu:", val)
	q.ImprimirLista()
	fmt.Println("Size:", q.Size())

	fmt.Println("\n>>> Teste DequeueRear")
	val, _ = q.DequeueRear()
	fmt.Println("Saiu:", val)
	q.ImprimirLista()
	fmt.Println("Size:", q.Size())

	val, _ = q.DequeueRear()
	fmt.Println("Saiu:", val)
	q.ImprimirLista()
	fmt.Println("Size:", q.Size())

	fmt.Println("\n>>> Esvaziando totalmente")
	val, _ = q.DequeueRear()
	fmt.Println("Saiu:", val)
	q.ImprimirLista()
	fmt.Println("IsEmpty:", q.IsEmpty(), "Size:", q.Size())

	fmt.Println("\n>>> Teste remoção em lista vazia")
	_, err := q.DequeueFront()
	if err != nil {
		fmt.Println("Erro esperado:", err)
	}
	_, err = q.DequeueRear()
	if err != nil {
		fmt.Println("Erro esperado:", err)
	}
}
