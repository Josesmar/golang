package main

import (
	"fmt"
)

func main() {
	tamanho := 40
	tarefas := make(chan int, tamanho)
	resultados := make(chan int, tamanho)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < tamanho; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < tamanho; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
