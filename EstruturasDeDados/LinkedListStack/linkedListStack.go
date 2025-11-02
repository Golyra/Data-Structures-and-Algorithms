package main

import (
	"fmt"
)

type IStack interface {
	Push(value int)
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
	Size() int
}

type no struct {
	val  int
	prox *no
}

type linkedListStack struct {
	top     *no
	tamanho int
}

func (lLS *linkedListStack) Peek() (int, error) {
	if lLS.tamanho == 0 {
		return -1, fmt.Errorf("A lista está vazia.")
	}
	return lLS.top.val, nil
}

func (lLS *linkedListStack) Size() int {
	return lLS.tamanho
}

func (lLS *linkedListStack) Push(value int) {
	novoNo := &no{val: value}
	if lLS.top == nil {
		lLS.top = novoNo
	} else {
		novoNo.prox = lLS.top
		lLS.top = novoNo
	}
	lLS.tamanho++
}

func (lLS *linkedListStack) Pop() (int, error) {
	aux := lLS.top
	aux_2 := lLS.top
	if lLS.tamanho == 0 {
		return -1, fmt.Errorf("A lista está vazia.")
	}
	lLS.top.val = 0
	lLS.top = aux.prox
	aux_2.prox = nil
	lLS.tamanho--
	return aux.val, nil
}

func (lLS *linkedListStack) IsEmpty() bool {
	if lLS.tamanho == 0 {
		return true
	} else {
		return false
	}
}

// Função pra facilitar visualização da lista
func (lLS *linkedListStack) Display() (string, error) {
	if lLS.top == nil {
		return "", fmt.Errorf("lista vazia")
	}

	aux := lLS.top
	result := ""

	for aux != nil {
		result += fmt.Sprintf("%d", aux.val)

		if aux.prox != nil {
			result += ", "
		}
		aux = aux.prox
	}

	return result, nil
}

// ============================= TESTES =============================
func main() {
	fmt.Println("=== Testando linkedListStack ===")

	stack := &linkedListStack{}

	// Teste inicial
	fmt.Println("\nA pilha está vazia?", stack.IsEmpty())
	fmt.Println("Tamanho inicial:", stack.Size())

	// Teste Push
	fmt.Println("\nInserindo elementos 10, 20, 30...")
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	str, _ := stack.Display()
	fmt.Println("Pilha:", str)
	fmt.Println("Tamanho:", stack.Size())

	// Teste Peek
	val, err := stack.Peek()
	if err == nil {
		fmt.Println("\nTopo da pilha (Peek):", val)
	} else {
		fmt.Println("Erro:", err)
	}

	// Teste Pop
	fmt.Println("\nRemovendo elementos com Pop...")
	rem, err := stack.Pop()
	if err == nil {
		fmt.Println("Elemento removido:", rem)
	}
	str, _ = stack.Display()
	fmt.Println("Pilha após Pop:", str)
	fmt.Println("Tamanho:", stack.Size())

	// Remover até esvaziar
	stack.Pop()
	stack.Pop()
	fmt.Println("\nEsvaziando a pilha...")
	fmt.Println("Está vazia?", stack.IsEmpty())
	fmt.Println("Tamanho:", stack.Size())

	// Teste Pop em pilha vazia
	fmt.Println("\nTentando Pop em pilha vazia...")
	_, err = stack.Pop()
	if err != nil {
		fmt.Println("Erro:", err)
	}

	// Teste Peek em pilha vazia
	fmt.Println("\nTentando Peek em pilha vazia...")
	_, err = stack.Peek()
	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Println("\nFim dos testes!")
}
