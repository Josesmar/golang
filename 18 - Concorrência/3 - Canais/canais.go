package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	// for {
	// 	mensagem, aberto := <-canal //Esperando que o canal recebe uma valor
	// 	if !aberto {                //Se o canal não estiver aberto eu saio desse loop
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal{
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal) //Fechando o canal
}
