package main

import "fmt"

//Defer significa que a função será adiata no último momento possível
//Pode ser usado para fechar conexão com banco de dados
func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado")            //Segundo executado
	fmt.Println("Entrando na função para verificar se o aluno está aprovado") //Primero executado

	media := (n1 + n2) / 2

	if media >= 6 {
		return true //Terceiro executado
	}
	return false //ou Terceiro executado
}

func main() {
	fmt.Println(alunoEstaAprovado(7, 8))

}
