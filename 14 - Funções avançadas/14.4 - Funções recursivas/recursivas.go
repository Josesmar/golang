package main

import "fmt"

// Função revursiva tem que ter um momento de parada, se não gera estouro de pilha
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	// Fibonacci
	// 1 1 2 3 5 8 13

	posicao := uint(8)
	fmt.Println("Soma", fibonacci(posicao))

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
