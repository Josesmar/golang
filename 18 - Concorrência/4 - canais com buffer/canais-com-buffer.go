package main

import "fmt"

func main() {
	canal := make(chan string, 2) //Capacidade de 2 buffer

	canal <- "Olá mundo!"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem, mensagem2)
}
