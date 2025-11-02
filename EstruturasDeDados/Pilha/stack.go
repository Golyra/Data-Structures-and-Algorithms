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

type ArrayStack struct {
	data     []int
	top      int
	capacity int
	length   int
}

func (s *ArrayStack) Init(size int) {
	s.data = make([]int, size)
	s.length = 0
	s.capacity = size
	s.top = -1
}

func (s *ArrayStack) Unstack() {
	if s.length == 0 {
		fmt.Println("Pilha vazia")
		return
	}

	fmt.Printf("[")
	for i := range s.length {
		fmt.Print(s.data[i])
		if i < s.length-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Println("]")
}

func (s *ArrayStack) doubleCapacity() {
	var newPack = make([]int, 2*s.capacity*2)
	for i := 0; i < s.length; i++ {
		newPack[i] = s.data[i]
	}
	s.data = newPack
	s.capacity = 2 * s.capacity
}

func (s *ArrayStack) Push(value int) {
	if s.length == s.capacity {
		s.doubleCapacity()
	}
	s.data[s.length] = value
	s.top = value
	s.length++
}

func (s *ArrayStack) Pop() (int, error) {
	if s.length == 0 {
		return -1, fmt.Errorf("A lista está vazia")
	}
	val := s.top
	s.data[s.length-1] = 0
	s.length--

	if s.length == 0 {
		s.top = -1
	} else {
		s.top = s.data[s.length-1]
	}

	return val, nil
}

func (s *ArrayStack) Peek() (int, error) {
	if s.top == -1 {
		return -1, fmt.Errorf("A lista está vazia")
	} else {
		return s.top, nil
	}
}

func (s *ArrayStack) IsEmpty() bool {
	if s.top == -1 {
		return true
	} else {
		return false
	}
}

func (s *ArrayStack) Size() int {
	return s.length
}

func isBalanced(expr string) bool {
	s := &ArrayStack{}
	s.Init(len(expr))

	for _, char := range expr {
		switch char {
		case '(':
			s.Push(1)
		case ')':
			if s.IsEmpty() {
				return false
			} else {
				s.Pop()
			}
		}
	}

	return s.IsEmpty()
}

// =================================Testes================================
func main() {
	stack := &ArrayStack{}
	stack.Init(5)

	// Testando Push
	fmt.Println("Adicionando elementos 10, 20, 30...")
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Printf("Estado da pilha:")
	stack.Unstack()

	// Testando Peek
	val, _ := stack.Peek()
	fmt.Println("Elemento no topo: ", val)

	// Testando IsEmpty
	isEmpty := stack.IsEmpty()
	fmt.Println("A pilha está vazia?", isEmpty)

	// Testando Size
	fmt.Println("Tamanho da pilha: ", stack.Size())

	// Testando Pop
	fmt.Println("\nFazendo Pop...")
	val, err := stack.Pop()
	if err == nil {
		fmt.Println("Elemento removido:", val) // esperado: 30
	} else {
		fmt.Println("Erro:", err)
	}
	fmt.Printf("Estado da pilha:")
	stack.Unstack()

	// Mais Pops
	val, _ = stack.Pop()
	fmt.Println("Removido:", val) // esperado: 20

	val, _ = stack.Pop()
	fmt.Println("Removido:", val) // esperado: 10

	// Pop em pilha vazia
	_, err = stack.Pop()
	if err != nil {
		fmt.Println("Erro esperado:", err)
	}

	// Teste de IsBalanced
	examples := []string{
		"()", "(a+b)", "((a+b)*c)", "(", ")(a+b)", "((a+b)",
	}

	for _, expr := range examples {
		fmt.Printf("%s -> %v\n", expr, isBalanced(expr))
	}
}
