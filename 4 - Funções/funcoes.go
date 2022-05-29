package main

import "fmt"

func somar(n1, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(2, 5)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("texto da função 1")
	resultadoSoma, pegarOutroRetorno := calculosMatematicos(10, 10) // usando o anderline é ignorado o valor de retorno da função
	println(resultadoSoma)
	fmt.Println(pegarOutroRetorno)
}
