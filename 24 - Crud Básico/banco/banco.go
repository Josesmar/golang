package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //import implicito - Driver de conexao com o MySQL
)

//Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local" //usuario:senha@/nomeBanco?configuracoes

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro //O valor zero de ponteiro é sempre nil
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
