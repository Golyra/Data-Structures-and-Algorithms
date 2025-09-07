package main

import (
	"fmt"
)

type no struct {
	val  int
	prox *no
}

type List interface {
	Size() int
	Add(val int)
	Remove(val int) (int, error)
	Get(index int) (int, error)
}

type linkedList struct {
	cabeca  *no
	tamanho int
}

func (lL *linkedList) Get(index int) (int, error) {
	aux := lL.cabeca
	for range index {
		aux = aux.prox
		if aux == nil {
			return -1, fmt.Errorf("índice fora do alcance da lista.")
		}
	}
	return aux.val, nil
}

func (lL *linkedList) Size() int {
	return lL.tamanho
}

func (lL *linkedList) Add(val int) {
	novoNo := &no{val: val}
	if lL.cabeca == nil {
		lL.cabeca = novoNo
	} else {
		aux := lL.cabeca
		for aux.prox != nil {
			aux = aux.prox
		}
		aux.prox = novoNo
	}
	lL.tamanho++
}

func (lL *linkedList) Remove(val int) (int, error) {
	aux := lL.cabeca
	aux_ant := lL.cabeca
	for aux.val != val {
		aux_ant = aux
		aux = aux.prox

		if aux == nil {
			return -1, fmt.Errorf("Elemento dado não existe na lista.")
		}
	}
	*aux_ant.prox = *aux.prox
	lL.tamanho--
	return 1, nil
}

func (lL *linkedList) Display() (string, error) {
	if lL.cabeca == nil {
		return "", fmt.Errorf("lista vazia")
	}

	aux := lL.cabeca
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

func main() {
	fmt.Println("=== Testando linkedList ===")

	list := &linkedList{}

	// Teste Add
	fmt.Println("\nAdicionando elementos 1..5")
	for i := 1; i <= 5; i++ {
		list.Add(i)
	}

	fmt.Println("Elementos na linked list: ")
	str, err := list.Display()
	fmt.Println("Lista: ", str)
	fmt.Println("Tamanho:", list.Size())

	// Teste Get
	val, err := list.Get(2)
	if err == nil {
		fmt.Println("\nElemento no índice 2:", val)
	} else {
		fmt.Println(err)
	}

	// Teste Remove
	fmt.Println("\nRemovendo o elemento 3...")
	rem, err := list.Remove(3)
	if err == nil {
		fmt.Println("Status: ", rem)
	} else {
		fmt.Println("Status: ", err)
	}
	str, err = list.Display()
	fmt.Println("Lista: ", str)
	fmt.Println("Tamanho:", list.Size())

	// Teste Remove de elemento inexistente
	fmt.Println("\nTentando remover 99...")
	_, err = list.Remove(99)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	// Teste Get fora do intervalo
	fmt.Println("\nTentando acessar índice 10...")
	_, err = list.Get(10)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Println("\nFim dos testes!")
}
