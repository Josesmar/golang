package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários"))
}

func main() {

	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	fmt.Printf("Escutando porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
