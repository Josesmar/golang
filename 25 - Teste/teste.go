package main

import "fmt"

func main() {
	fmt.Println("Fatorial")
	fmt.Println(fatorial(5))

	fmt.Println("Fibonacci")
	posicao := int(12) //Declarou o valor a ser somado
	// fmt.Println("Soma", fibo(posicao))

	for i := int(1); i <= posicao; i++ {
		fmt.Println(fibo(i))
	}

}

func fatorial(numero int) int {

	if numero == 0 {
		return 1
	} else {
		return numero * fatorial(numero-1)
	}
}

func fibo(valor int) int {
	if valor <= 1 {
		return valor
	} else {
		return fibo(valor-2) + fibo(valor-1)
	}
}
