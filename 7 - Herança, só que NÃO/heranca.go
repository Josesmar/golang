package main

import "fmt"

type pessoa struct {
	nome      string
	sobreNome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	fmt.Println("Herança só que NÃO")

	p1 := pessoa{"Josesmar","Santos", 20, 160}
	fmt.Println(p1)

	e1 := estudante{p1, "Ciência da computação", "UFMG"}
	fmt.Println(e1.altura)

}
