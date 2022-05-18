package main

import "fmt"

//Passando parâmetro por cópia
func inverterSinal(numero int) int {
	return numero * -1
}

//Não precisa de retorno, pois a alteração é feita diretamente na variavel declarada no main
//Passando uma referência para essa função
func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {

	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	//usando & comercial é inidicado que vai usar o endereço de memória
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}
