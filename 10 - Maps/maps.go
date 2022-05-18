package main

import "fmt"

func main() {
	fmt.Println(" ---  Maps --- ")

	// Maps - Dentro dos couchetes é o tipo da chave, fora do couchetes é o tipo do valor
	// O Map é mais rigido quanto ao tipo de dados, aceitando apenas os tipos definido
	// A chave e valor precisam ser do mesmo tipo
	// Não é possível acessar a variavel usando o ponto
	// Para acessar só colocando o nome da variavel da seguinte forma >> usuario["nome"]
	usuario := map[string]string{
		"nome":      "Josesmar",
		"sobrenome": "Santos",
	}

	fmt.Println(usuario)

	// Usando map dentro de map
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Josesmar",
			"ultimo":   "Santos",
		},
		"curso": {
			"curso":  "Engenharia",
			"campos": "Campos 01",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	// Deletando uma chave do map
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"signo": "Virgem",
	}
	fmt.Println(usuario2)
}
