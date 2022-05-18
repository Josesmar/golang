package main

import "fmt"

func main() {

	fmt.Println("*Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println("Antes")
	fmt.Println("Varivel1: ", variavel1, "Variavel2: " , variavel2)
	
	variavel1++;

	fmt.Println("Depois")
	fmt.Println("Varivel1: ", variavel1, "Variavel2: " , variavel2)

	//PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3  //Adicionando referência de memória da veriavel3 na variavel ponteiro

	fmt.Println("-----Ponteiro-----")
	fmt.Println(variavel3, ponteiro) //Desreferenciando o endereço e pegando o valor
	fmt.Println(variavel3, *ponteiro) //Desreferenciando o endereço e pegando o valor

	fmt.Println("Alterando valor da veriavel3 e mantendo a referência do ponteiro")
	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)
}

