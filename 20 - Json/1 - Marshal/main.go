package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

/*
	json.Marshal(v any) //Converter mapper ou struct para um Json
	json.Unmarshal(data []byte, v any) //Converter um Json em Mapper ou Struct
*/

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cachorroEmJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJson)                  //Devolveu um slice de bytes uint8
	fmt.Println(bytes.NewBuffer(cachorroEmJson)) //Retorna um Json

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}
	cachorro2EmJson, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJson)                  //Devolveu um slice de bytes uint8
	fmt.Println(bytes.NewBuffer(cachorro2EmJson)) //Retorna um Json

}
