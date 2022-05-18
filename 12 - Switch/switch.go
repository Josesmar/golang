package main

import "fmt"

//Switch no GO não usa o break

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
		//fallthrough joga o resuldao para o próximo valor
		//fallthrough
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sabado"
	default:
		return "Numero inválido"
	}

}

func main() {
	fmt.Println(" ---- Switch ----")
	dia := diaDaSemana(8)
	fmt.Println(dia)
	fmt.Println(" ---- Dia 02 ----")
	dia2 := diaDaSemana(7)
	fmt.Println(dia2)

}
