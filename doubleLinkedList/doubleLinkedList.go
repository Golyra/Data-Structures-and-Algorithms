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

type DoubleLinkedList struct {
	head   *Node
	length int
	tail   *Node
}

type Node struct {
	prev  *Node
	value int
	next  *Node
}

func (DLL *DoubleLinkedList) ImprimirLista() {
	// usamos o aux para percorrer a lista
	if (DLL.head) != nil {
		aux := (DLL.head)
		// navega partindo da cabeça até chegar NULL
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

func (DLL *DoubleLinkedList) Add(e int) {
	newNode := &Node{value: e, prev: DLL.tail}
	if DLL.head == nil {
		DLL.head = newNode
		DLL.tail = newNode
	} else {
		DLL.tail.next = newNode
		DLL.tail = newNode
	}
	DLL.length++
}

func (DLL *DoubleLinkedList) Push(e int) {
	newNode := &Node{value: e, next: DLL.head}
	if DLL.head == nil {
		DLL.head = newNode
		DLL.tail = newNode
	} else {
		DLL.head.prev = newNode
		DLL.head = newNode
	}
	DLL.length++
}

func (DLL *DoubleLinkedList) InsertOnIndex(e int, index int) {
	newNode := &Node{value: e}
	mid := DLL.length / 2
	if index <= mid {
		aux := DLL.head
		for range index {
			aux = aux.next
		}
		newNode.prev = aux.prev
		aux.prev = newNode
		newNode.next = aux
		newNode.prev.next = newNode
	} else {
		aux := DLL.tail
		for i := DLL.length; i > index+1; i-- {
			aux = aux.prev
		}
		newNode.prev = aux.prev
		aux.prev = newNode
		newNode.next = aux
		newNode.prev.next = newNode
	}
	DLL.length++
}

func (DLL *DoubleLinkedList) Get(index int) (int, error) {
	mid := DLL.length / 2
	if index > DLL.length || index < 0 {
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
		for i := DLL.length; i > index+1; i-- {
			aux = aux.prev
		}
		return aux.value, nil
	}
}

func main() {
	fmt.Println("=== Testando Double Linked List ===")
	fmt.Println("")

	list := &DoubleLinkedList{}

	fmt.Println("- Testando função de impressão")
	list.ImprimirLista()
	fmt.Println("")

	fmt.Println("- Testando adição no fim da lista")
	list.Add(1983)
	list.Add(1987)
	list.Add(2023)
	list.Add(2035)
	list.ImprimirLista()
	fmt.Println("Tamanho da lista: ", list.length)
	fmt.Println("")

	fmt.Println("- Testando adição no início da lista")
	list.Push(1979)
	list.ImprimirLista()
	fmt.Println("Tamanho da lista: ", list.length)
	fmt.Println("")

	fmt.Println("- Testando inserção em índice")
	list.InsertOnIndex(1985, 2)
	list.ImprimirLista()
	fmt.Println("Tamanho da lista: ", list.length)
	fmt.Println("")

	fmt.Println("- Testando mais uma inserção em índice")
	list.InsertOnIndex(1993, 4)
	list.ImprimirLista()
	fmt.Println("Tamanho da lista: ", list.length)
	fmt.Println("")

	fmt.Println("- Testando busca por índice")
	Ans, err := list.Get(3)
	if err == nil {
		fmt.Printf("Is that the bite of %d ", Ans)
	} else {
		fmt.Printf(err.Error())
	}
	fmt.Println("")
	Ans, err = list.Get(2)
	if err == nil {
		fmt.Printf("Missing Children Incident of %d", Ans)
	} else {
		fmt.Printf(err.Error())
	}
}
