package config

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	username  = "root"
	password  = "widnyana99"
	hostname  = "127.0.0.1:3306"
	dbName    = "insert_million"
	maxConn   = 100
	maxIdle   = 5
	maxWorker = 100
)

var connections = fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)

func ConnectDB() (*sql.DB, error) {
	log.Println("Start connections to database......")

	db, err := sql.Open("mysql", connections)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdle)

	return db, nil
}
