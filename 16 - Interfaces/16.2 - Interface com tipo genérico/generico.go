package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {

	generica("String")
	generica(1)
	generica(true)

	//é um tipo de interface genérica
	fmt.Println(1, 2, "String", false, true, float64(1231321))

}
