package main

import "fmt"

func escrever() {
	fmt.Println("Escrevendo...")
}

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
	fmt.Println()
}

func (u usuario) maiorIdade() bool {
	return u.idade > 18

}

//Altera o valor da veriaval
func (u *usuario) fazerAniversario() {
	u.idade++ //só chama normalmente sem desreferenciar
}

func main() {
	usuario1 := usuario{"Josesmar", 20}
	if usuario1.maiorIdade() {
		usuario1.salvar()
	} else {
		fmt.Println("Necessário responsável para salvar usuário")
	}

	usuario2 := usuario{"Pamela", 17}

	if usuario2.maiorIdade() {
		usuario2.salvar()
	} else {
		fmt.Printf("Necessário responsável para salvar usuário %s\n", usuario2.nome)
	}

	usuario2.fazerAniversario()
	fmt.Println(usuario1.idade)

}
