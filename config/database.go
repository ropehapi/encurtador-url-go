package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetConexao() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_encurtador_url_go")

	if err != nil {
		panic(err.Error())
	}

	return db
}
