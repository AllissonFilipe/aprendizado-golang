package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //meu arquivo nao vai usar o drive mas sim o pacote sql pra identificar qual é o sgbd
)

func main() {
	stringConexao := "root:@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil { // tou reaproveitando a variavel de erro
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		log.Fatal(erro)
	}

	defer linhas.Close()

	fmt.Println(linhas)
}
