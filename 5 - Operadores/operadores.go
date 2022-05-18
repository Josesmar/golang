package main

import "fmt"

func main() {
	//ARITMETICOS - + / * %
	soma := 1 + 3
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//FORTEMENTE TIPADO não pode somar tipos de int diferentes
	var numero1 int16 = 20
	var numero2 int16 = 25
	var soma2 int16 = numero1 + numero2
	fmt.Println(soma2)
	//FIM DOS ARITIMÉTICOS

	//ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "string 2"
	fmt.Println(variavel1, variavel2)
	//FIM DOS OPERADORES DE ATRIBUICAO

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	//FIM DOS RELACIONAIS

	//OPERADORES LÓGICOS
	fmt.Println("-------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || !falso)
	//FIM DOS OPERADORES LÓGICOS

	//OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15 // numero = numero -20
	numero -= 20 // numero = numero * 3
	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)
    //FIM DOS OPERADORES UNIÁRIO

	//OPERADOR TERNÁRIO	
	var texto string
	if(numero > 5){
		texto = "Maior que 5"
	} else{
		texto = "Menor que 5"
	}
	fmt.Println(texto)
	 
}
