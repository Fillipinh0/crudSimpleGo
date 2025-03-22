package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func ConectaBanco() *sql.DB {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "1234",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "db_trabalho",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Erro ao abrir conexão com o banco de dados:", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", pingErr)
	}

	fmt.Println("Conexão realizada com sucesso!")
	return db
}
