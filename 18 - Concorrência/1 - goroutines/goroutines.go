package main

//Concorrência != Paralelismo
import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá mundo!") // goroutine (Executa essa função, mas não espera terminar a execução para seguir com as outras execuções)
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
