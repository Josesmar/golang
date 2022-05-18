package main

import "fmt"

func main() {

	//A função anonima dentro do main, estou dizendo para GO...oh esse é uma função anonima
	//Após a declaração dela usar os parênteses no final da chave da função para dizer ao GO, ohh agora execute ()
	retorno := func(texto string) string {
		return fmt.Sprintf("recebido -> %s", texto)
	}("Passando o valor do parâmetro da função anônima")

	fmt.Println(retorno)

}
