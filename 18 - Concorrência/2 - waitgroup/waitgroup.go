package main

//Concorrência != Paralelismo
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup //Grupo de espera

	waitGroup.Add(4) //Informando que tem duas goroutinas que fazem parte desse grupo de espera

	//Função anonima
	go func() {
		escrever("Olá mundo!")
		waitGroup.Done() //Tirar uma goroutine (-1)
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutine 3!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutine 4!")
		waitGroup.Done()
	}()

	waitGroup.Wait() //Esperar a contagem da goroutine chegarem em zero (Espera que as funções terminei de serem executadas)

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
