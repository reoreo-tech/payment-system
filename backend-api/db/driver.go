package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func init() {
	user := os.Getenv("MYSQL_USER")
	// host := os.Getenv("MYSQL_HOST")
	pass := os.Getenv("MYSQL_PASSWORD")
	name := os.Getenv("MYSQL_DATABASE")
	// port := os.Getenv("MYSQL_PORT")

	dbconf := user + ":" + pass + "@/" + name
	conn, err := sql.Open("mysql", dbconf)
	if err != nil {
		panic(err.Error)
	}
	Conn = conn
}
