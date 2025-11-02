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

type DoubleLinkedList struct {
	head   *Node
	length int
	tail   *Node
}

type Node struct {
	prev *Node
	val  int
	next *Node
}

// Pra facilitar a visualização dos testes
func (DLL *DoubleLinkedList) ImprimirLista() {

	if (DLL.head) != nil {
		aux := (DLL.head)
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

func (DLL *DoubleLinkedList) Add(value int) {
	newNode := &Node{val: value, prev: DLL.tail}
	if DLL.head == nil {
		DLL.head = newNode
		DLL.tail = newNode
	} else {
		DLL.tail.next = newNode
		newNode.prev = DLL.tail
		DLL.tail = newNode
	}
	DLL.length++
}

func (DLL *DoubleLinkedList) AddOnIndex(value int, index int) {
	newNode := &Node{val: value}
	mid := DLL.length / 2
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
	} else if index == DLL.length {
		newNode.prev = DLL.tail
		DLL.tail.next = newNode
		DLL.tail = newNode
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
		return aux.val, nil
	} else {
		aux := DLL.tail
		for i := DLL.length; i > index+1; i-- {
			aux = aux.prev
		}
		return aux.val, nil
	}
}

func (DLL *DoubleLinkedList) Size() int {
	return DLL.length
}

func (DLL *DoubleLinkedList) Set(value int, index int) error {
	if index < 0 || index >= DLL.length {
		return fmt.Errorf("índice fora de alcance")
	}

	var aux *Node
	mid := DLL.length / 2

	if index <= mid {
		aux = DLL.head
		for range index {
			aux = aux.next
		}
	} else {
		aux = DLL.tail
		for i := DLL.length - 1; i > index; i-- {
			aux = aux.prev
		}
	}

	aux.val = value
	return nil
}

func (DLL *DoubleLinkedList) RemoveOnIndex(index int) error {
	if index < 0 || index >= DLL.length {
		return fmt.Errorf("índice fora de alcance")
	}

	var aux *Node
	mid := DLL.length / 2

	if index <= mid {
		aux = DLL.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
	} else {
		aux = DLL.tail
		for i := DLL.length - 1; i > index; i-- {
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

	DLL.length--
	return nil
}

// =========================================== TESTES ===========================================
func main() {
	fmt.Println("=== Testando DoubleLinkedList ===")

	list := &DoubleLinkedList{}

	// -------------------- Teste Add --------------------
	fmt.Println("\nAdicionando elementos 1, 2, 3, 4, 5")
	for i := 1; i <= 5; i++ {
		list.Add(i)
	}
	fmt.Print("Lista após Add: ")
	list.ImprimirLista()
	fmt.Println("Tamanho:", list.Size())

	// -------------------- Teste Get --------------------
	val, err := list.Get(2) // índice 2 = 3
	if err != nil {
		fmt.Println("Erro Get:", err)
	} else {
		fmt.Println("\nElemento no índice 2 (Get):", val)
	}

	// -------------------- Teste Set --------------------
	fmt.Println("\nAlterando índice 2 para 99 (Set)")
	if err := list.Set(99, 2); err != nil {
		fmt.Println("Erro Set:", err)
	} else {
		fmt.Print("Lista após Set: ")
		list.ImprimirLista()
	}

	// -------------------- Teste AddOnIndex --------------------
	fmt.Println("\nInserindo 77 no índice 1 (AddOnIndex)")
	list.AddOnIndex(77, 1)
	fmt.Print("Lista após AddOnIndex: ")
	list.ImprimirLista()
	fmt.Println("Tamanho:", list.Size())

	fmt.Println("\nInserindo 88 no índice 0 (AddOnIndex na cabeça)")
	list.AddOnIndex(88, 0)
	fmt.Print("Lista após AddOnIndex: ")
	list.ImprimirLista()
	fmt.Println("Tamanho:", list.Size())

	fmt.Println("\nInserindo 99 no final (AddOnIndex na cauda)")
	list.AddOnIndex(99, list.Size())
	fmt.Print("Lista após AddOnIndex: ")
	list.ImprimirLista()
	fmt.Println("Tamanho:", list.Size())

	// -------------------- Teste RemoveOnIndex --------------------
	fmt.Println("\nRemovendo índice 3 (RemoveOnIndex)")
	if err := list.RemoveOnIndex(3); err != nil {
		fmt.Println("Erro RemoveOnIndex:", err)
	} else {
		fmt.Print("Lista após remoção: ")
		list.ImprimirLista()
		fmt.Println("Tamanho:", list.Size())
	}

	fmt.Println("\nRemovendo índice 0 (RemoveOnIndex na cabeça)")
	if err := list.RemoveOnIndex(0); err != nil {
		fmt.Println("Erro RemoveOnIndex:", err)
	} else {
		fmt.Print("Lista após remoção: ")
		list.ImprimirLista()
		fmt.Println("Tamanho:", list.Size())
	}

	fmt.Println("\nRemovendo último índice (RemoveOnIndex na cauda)")
	if err := list.RemoveOnIndex(list.Size() - 1); err != nil {
		fmt.Println("Erro RemoveOnIndex:", err)
	} else {
		fmt.Print("Lista após remoção: ")
		list.ImprimirLista()
		fmt.Println("Tamanho:", list.Size())
	}

	// -------------------- Testes de Get e Set fora do intervalo --------------------
	fmt.Println("\nTentando acessar índice inválido (Get)")
	if _, err := list.Get(100); err != nil {
		fmt.Println("Erro esperado:", err)
	}

	fmt.Println("\nTentando alterar índice inválido (Set)")
	if err := list.Set(123, 100); err != nil {
		fmt.Println("Erro esperado:", err)
	}

	fmt.Println("\nFim dos testes!")
}
