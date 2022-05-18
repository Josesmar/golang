package main

import "fmt"

var n int

//Função init sempre vai ser executada antes do main por arquivo
func init() {
	n = 10
	fmt.Println("Executando a função init")
}

func main() {

	fmt.Println("Função main sendo executada")
	fmt.Println(n)

}
