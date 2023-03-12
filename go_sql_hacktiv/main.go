package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

var (
	username    = "root"
	password    = "widnyana99"
	hostname    = "127.0.0.1:3306"
	dbName      = "kampus_merdeka"
	connections = fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", connections)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	sqlQuery()
	sqlQueryRow()
	sqlPrepare()
	sqlExec()
}
