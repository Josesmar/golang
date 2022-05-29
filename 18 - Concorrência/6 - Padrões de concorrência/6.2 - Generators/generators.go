package main

/* User o generator para esconder a complexidade
 */

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá mundo!") //Canal que só recebe

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
