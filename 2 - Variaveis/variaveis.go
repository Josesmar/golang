package main

import "fmt"

func main() {
	var variavel1 string = "Variavel1"
	variavel2 := "Variavel 2"
	
	fmt.Println(variavel1)
	fmt.Println(variavel2)


	var(
		variavel3 string = "balbla"
		variavel4 string = "vaadfasdf"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel5", "Varivel6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	//trocar valores
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

	

}