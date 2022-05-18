package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {

	 soma, subtracao := calculosMatematicos(2, 1)

	fmt.Println("soma: ", soma)
	fmt.Println("subtracao: ", subtracao)

}
