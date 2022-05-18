package main

import (
	"errors"
	"fmt"
)

func main() {
	//quantidade de bits
	//int8, int16, int32, int64

	var numero8 int8 = 100
	var numero16 int16 = 10000
	var numero32 int32 = 1000000000
	var numero64 int64 = 1000000000000000000
	var numero int = 1000000000000000000
	fmt.Println(numero8)
	fmt.Println(numero16)
	fmt.Println(numero32)
	fmt.Println(numero64)
	fmt.Println(numero)

	//unsygned int (int sem sinal)
	var numero2 uint32 = 10
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	//NUMEROS REAIS
	//float32, float64
	var numeroReal1 float32 = 132.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1300000002.45
	fmt.Println(numeroReal2)

	numeroReal3 := 13000.00
	fmt.Println(numeroReal3)

	//FIM NUMEROS REAIS

	//STRING
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texte2"
	fmt.Println(str2)

	char := 'B' //NÚMERO DA TABELA ASCI
	fmt.Println(char)

	//FIM STRING

	var texto int16
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
