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

type no struct {
	val  int
	prox *no
}

type linkedListQueue struct {
	cabeca  *no
	tamanho int
}

func (lLD *linkedListQueue) Enqueue(value int) {
	novoNo := &no{val: value}
	if lLD.cabeca == nil {
		lLD.cabeca = novoNo
	} else {
		aux := lLD.cabeca
		for aux.prox != nil {
			aux = aux.prox
		}
		aux.prox = novoNo
	}
	lLD.tamanho++
}

func (lLD *linkedListQueue) Dequeue() (int, error) {
	if lLD.tamanho == 0 {
		return -1, fmt.Errorf("A lista está vazia")
	} else {
		a := lLD.cabeca.val
		aux := lLD.cabeca

		lLD.cabeca = lLD.cabeca.prox
		lLD.tamanho--

		aux.val = 0
		aux.prox = nil

		return a, nil
	}

}

func (lLD *linkedListQueue) Front() (int, error) {
	if lLD.tamanho <= 0 {
		return -1, fmt.Errorf("A lista está vazia.")
	} else {
		return lLD.cabeca.val, nil
	}
}

func (lLD *linkedListQueue) IsEmpty() bool {
	if lLD.tamanho <= 0 {
		return true
	} else {
		return false
	}
}

func (lLD *linkedListQueue) Size() int {
	return lLD.tamanho
}

// Função pra facilitar visualização da lista
func (lLD *linkedListQueue) Display() (string, error) {
	if lLD.cabeca == nil {
		return "", fmt.Errorf("lista vazia")
	}

	aux := lLD.cabeca
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
	fmt.Println("=== Testando linkedListQueue ===")

	q := &linkedListQueue{}

	// Teste Enqueue
	fmt.Println("\nAdicionando elementos 10, 20, 30")
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	str, _ := q.Display()
	fmt.Println("Fila:", str)
	fmt.Println("Tamanho:", q.Size())
	fmt.Println("Está vazia?", q.IsEmpty())

	// Teste Front
	front, err := q.Front()
	if err == nil {
		fmt.Println("Elemento da frente:", front)
	} else {
		fmt.Println("Erro:", err)
	}

	// Teste Dequeue
	fmt.Println("\nRemovendo elementos da fila")
	for !q.IsEmpty() {
		val, err := q.Dequeue()
		if err == nil {
			fmt.Println("Dequeued:", val)
			str, _ := q.Display()
			fmt.Println("Fila restante:", str)
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
