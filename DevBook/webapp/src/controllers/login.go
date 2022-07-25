package controllers

import (
	"net/http"
)

// CarregarTelaDeLogin vai denderizar tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	// utils.ExecutarTemplate(w, "login.html", nil)
	w.Write([]byte("Tela de login"))
}
