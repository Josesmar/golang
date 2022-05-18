package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	//COLOCANDO VALOR FIXO NO ARRAY
	fmt.Println("--- Valor fixo do array -- ")
	array2 := [5]string{"posição 1", "posição 2", "posição 3", "posição 4", "posição 5"}
	fmt.Println(array2)

	//NÃO É MUITO UTILIZADO
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//MAIS UTILIZADO
	fmt.Println("--- Incluindo valores no array por Slice -- ")
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18}
	fmt.Println(slice)

	slice = append(slice, 15)
	fmt.Println(slice)

	//TypeOf retorno o tipo da variavel
	fmt.Println("--- Tipos dos arrays -- ")
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	fmt.Println("--- Pegando uma parte do array (Pega indice 1 e 2) -- ")
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	// Arrays internos
	fmt.Println("----------")
	fmt.Println("Arrays internos")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	//Quando o GO ver que a sua capacidade do slice vai estourar, ele cria um outro array e dobra o valor dele
	//Eu inha um array de capacidade 11 e após a inclusão dos valores abaixo foi para 12 e o GO criou um novo array de 24 posições
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	fmt.Println("Slice sem capacidade máxima definida")
	slice4 := make([]float32, 5)
	fmt.Println(slice4)

	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	//Array com tamanho fixo
	//Slice sem tamanho fixo
}
