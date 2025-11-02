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

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	prev  *Node
	value int
	next  *Node
}

// Pra facilitar a visualização dos testes
func (DLL *DoublyLinkedList) ImprimirLista() {

	if (DLL.head) != nil {
		aux := (DLL.head)
		fmt.Printf("[")

		for aux != nil {
			fmt.Printf("%d", aux.value)
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

func (list *DoublyLinkedList) Reverse() {
	var ant *Node = nil
	var aux *Node = list.head
	var nxt *Node = nil

	for aux != nil {
		nxt = aux.next
		aux.next = ant
		aux.prev = nxt
		ant = aux
		aux = nxt
	}

	list.tail = list.head
	list.head = ant
}

func (DLL *DoublyLinkedList) Add(value int) {
	newNode := &Node{value: value, prev: DLL.tail}
	if DLL.head == nil {
		DLL.head = newNode
		DLL.tail = newNode
	} else {
		DLL.tail.next = newNode
		newNode.prev = DLL.tail
		DLL.tail = newNode
	}
	DLL.size++
}

func (DLL *DoublyLinkedList) AddOnIndex(value int, index int) {
	newNode := &Node{value: value}
	mid := DLL.size / 2
	if index == 0 {
		newNode.next = DLL.head
		DLL.head = newNode
	} else if index <= mid {
		aux := DLL.head
		for range index {
			aux = aux.next
		}
		newNode.prev = aux.prev
		aux.prev = newNode
		newNode.next = aux
		newNode.prev.next = newNode
	} else if index == DLL.size {
		newNode.prev = DLL.tail
		DLL.tail.next = newNode
		DLL.tail = newNode
	} else {
		aux := DLL.tail
		for i := DLL.size; i > index+1; i-- {
			aux = aux.prev
		}
		newNode.prev = aux.prev
		aux.prev = newNode
		newNode.next = aux
		newNode.prev.next = newNode
	}
	DLL.size++
}

func (DLL *DoublyLinkedList) Get(index int) (int, error) {
	mid := DLL.size / 2
	if index > DLL.size || index < 0 {
		return -1, fmt.Errorf("Índice fora de alcance")
	}
	if index <= mid {
		aux := DLL.head
		for range index {
			aux = aux.next
		}
		return aux.value, nil
	} else {
		aux := DLL.tail
		for i := DLL.size; i > index+1; i-- {
			aux = aux.prev
		}
		return aux.value, nil
	}
}

func (DLL *DoublyLinkedList) Size() int {
	return DLL.size
}

func (DLL *DoublyLinkedList) Set(value int, index int) error {
	if index < 0 || index >= DLL.size {
		return fmt.Errorf("índice fora de alcance")
	}

	var aux *Node
	mid := DLL.size / 2

	if index <= mid {
		aux = DLL.head
		for range index {
			aux = aux.next
		}
	} else {
		aux = DLL.tail
		for i := DLL.size - 1; i > index; i-- {
			aux = aux.prev
		}
	}

	aux.value = value
	return nil
}

func (DLL *DoublyLinkedList) RemoveOnIndex(index int) error {
	if index < 0 || index >= DLL.size {
		return fmt.Errorf("índice fora de alcance")
	}

	var aux *Node
	mid := DLL.size / 2

	if index <= mid {
		aux = DLL.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
	} else {
		aux = DLL.tail
		for i := DLL.size - 1; i > index; i-- {
			aux = aux.prev
		}
	}

	// Ajusta os ponteiros prev e next
	if aux.prev != nil {
		aux.prev.next = aux.next
	} else {
		// Removendo a cabeça
		DLL.head = aux.next
	}
	if aux.next != nil {
		aux.next.prev = aux.prev
	} else {
		// Removendo a cauda
		DLL.tail = aux.prev
	}

	DLL.size--
	return nil
}

// =========================================== TESTES ===========================================
func main() {
	list := &DoublyLinkedList{}

	// Adicionando elementos
	for i := 1; i <= 5; i++ {
		list.Add(i)
	}
	fmt.Print("Lista original: ")
	list.ImprimirLista()

	// Revertendo
	list.Reverse()
	fmt.Print("Lista invertida: ")
	list.ImprimirLista()

	// Esperado: 5 4 3 2 1
}
