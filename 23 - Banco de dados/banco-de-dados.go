package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //import implicito
)

func main() {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local" //String de conexão mysql
	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {		
		log.Fatal(erro)
	}

	defer db.Close() //Antes da função main terminar ele vai fechar a conexão com o banco

	if erro = db.Ping(); erro != nil { //Para testar a conexão deve usar db.Ping()		
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	linhas, erro:= db.Query("select * from usuarios")
	if erro != nil {		
		log.Fatal(erro)
	}

	defer linhas.Close()

	fmt.Println(linhas)
}
