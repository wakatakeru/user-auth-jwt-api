package infrastructure

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	config := &mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		DBName:               os.Getenv("DB_NAME"),
		Addr:                 os.Getenv("DB_ADDR"),
		Net:                  "tcp",
		AllowNativePasswords: true,
	}

	conn, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Panic(err)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
