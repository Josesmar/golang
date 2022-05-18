package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--- Loops ----")
	total := 10
	i := 0

	for i < total {
		i++
		fmt.Println("Incrementando i", i)
		time.Sleep(time.Second)
		//fmt.Println("Contando de ", i , "a ", total)
	}

	//Contando de 2 em 2
	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"Josesmar", "Pamela", "Maura"}

	//Simbolo underline oculta o valor da variavel
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	//Simbolo underline oculta o valor da variavel
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	// interar por string é necessário fazer um cast string(letra), caso contrário o GO retornará o código da tabela ASCI
	for indice, letra := range "Palavra" {
		fmt.Println(indice, string(letra))
	}

	// Iterando um map
	usuario := map[string]string{
		"nome":      "Josesmar",
		"sobreNome": "Santos",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//Não é possivel fazer range em struct
	// type usuarioStruct struct {
	// 	nome      string
	// 	sobrenome string
	// }

	// usuario2 := usuarioStruct{"Zé", "Mané"}

	// for chave, valor := range usuario2 {
	// 	fmt.Println(chave, valor)
	// }

	
}
