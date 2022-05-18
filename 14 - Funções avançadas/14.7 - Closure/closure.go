package main

import "fmt"

//As Closure são funções que referênciam varráveis que estão fora do seu corpo

func closure() func() {
	texto := "Dentro da função clousure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {

	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
