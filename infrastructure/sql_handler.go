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

type SqlResult struct {
	Result sql.Result
}

type SqlRow struct {
	Rows *sql.Rows
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

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (SqlResult, error) {
	sqlResult := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return SqlResult{}, err
	}
	sqlResult.Result = result
	return sqlResult, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (SqlRow, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return SqlRow{}, err
	}
	row := SqlRow{}
	row.Rows = rows
	return row, nil
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
