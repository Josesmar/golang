package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.CarregarTemplate()
	r := router.Gerar()

	localhost := "127.0.0.1"
	porta := ":4000"
	fmt.Println("Escutando na porta ", porta)
	log.Fatal(http.ListenAndServe(localhost+porta, r))
}
