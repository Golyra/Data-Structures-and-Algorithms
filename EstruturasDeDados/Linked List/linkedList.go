package main

import (
	"fmt"
)

type no struct {
	val  int
	prox *no
}

type IList interface {
	Add(value int)
	AddOnIndex(value int, index int) error
	RemoveOnIndex(index int) error
	Get(index int) (int, error)
	Set(value int, index int) error
	Size() int
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

func (lL *linkedList) Set(value int, index int) error {
	if index < 0 || index >= lL.tamanho {
		return fmt.Errorf("índice fora do alcance da lista")
	}
	aux := lL.cabeca
	for i := 0; i < index; i++ {
		aux = aux.prox
	}
	aux.val = value
	return nil
}

func (lL *linkedList) Size() int {
	return lL.tamanho
}

func (lL *linkedList) Add(value int) {
	novoNo := &no{val: value}
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

func (lL *linkedList) AddOnIndex(value int, index int) error {
	if index < 0 || index > lL.tamanho {
		return fmt.Errorf("índice fora do alcance da lista")
	}
	novoNo := &no{val: value}

	if index == 0 {
		novoNo.prox = lL.cabeca
		lL.cabeca = novoNo
		lL.tamanho++
		return nil
	}
	aux := lL.cabeca
	for i := 0; i < index-1; i++ {
		aux = aux.prox
	}

	novoNo.prox = aux.prox
	aux.prox = novoNo
	lL.tamanho++
	return nil
}

func (lL *linkedList) RemoveOnIndex(index int) (int, error) {
	aux := lL.cabeca
	aux_ant := lL.cabeca
	if lL.tamanho == 0 {
		return -1, fmt.Errorf("A lista está vazia.")
	}
	if index > lL.tamanho || index < 0 {
		return -1, fmt.Errorf("índice não existe na lista")
	}
	for range index {
		aux_ant = aux
		aux = aux.prox
	}
	*aux_ant.prox = *aux.prox
	lL.tamanho--
	return 1, nil
}

// Função pra facilitar visualização da lista
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

// ============================= TESTES =============================
func main() {
	fmt.Println("=== Testando linkedList ===")

	list := &linkedList{}

	// Teste Add
	fmt.Println("\nAdicionando elementos 1..5")
	for i := 1; i <= 5; i++ {
		list.Add(i)
	}

	str, _ := list.Display()
	fmt.Println("Lista:", str)
	fmt.Println("Tamanho:", list.Size())

	// Teste Get
	val, err := list.Get(2)
	if err == nil {
		fmt.Println("\nElemento no índice 2:", val)
	} else {
		fmt.Println("Erro:", err)
	}

	// Teste Set
	fmt.Println("\nAlterando índice 2 para 99...")
	if err := list.Set(99, 2); err != nil {
		fmt.Println("Erro:", err)
	} else {
		str, _ := list.Display()
		fmt.Println("Lista após Set:", str)
	}

	// Teste AddOnIndex
	fmt.Println("\nInserindo 15 no índice 1...")
	if err := list.AddOnIndex(15, 1); err != nil {
		fmt.Println("Erro:", err)
	} else {
		str, _ := list.Display()
		fmt.Println("Lista após AddOnIndex:", str)
	}

	// Teste RemoveOnIndex
	fmt.Println("\nRemovendo no índice 3...")
	rem, err := list.RemoveOnIndex(3)
	if err == nil {
		fmt.Println("Elemento removido:", rem)
		str, _ = list.Display()
		fmt.Println("Lista após remoção:", str)
		fmt.Println("Tamanho:", list.Size())
	} else {
		fmt.Println("Erro:", err)
	}

	// Teste RemoveOnIndex em índice inexistente
	fmt.Println("\nTentando remover no índice 99...")
	_, err = list.RemoveOnIndex(99)
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
