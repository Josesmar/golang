package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// Carregar templates insere os tempplates html na variavel templates
func CarregarTemplate() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecutarTemplate renderiza uma página html na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)

}
