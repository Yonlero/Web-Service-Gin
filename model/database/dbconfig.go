package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	POSTGRESDRIVER = "postgres"
	USER           = "root"
	HOST           = "localhost"
	PORT           = "5432"
	PASSWORD       = "123"
	DBNAME         = "Db_Gin"
)

type (
	Database  struct{}
	IDatabase interface {
		OpenConnection() *sql.DB
		CloseConnection(db *sql.DB)
	}
)

var db *sql.DB
var err error
var datasourceName = fmt.Sprintf("host=%s "+
	"port=%s "+
	"user=%s "+
	"password=%s "+
	"dbname=%s "+
	"sslmode=disable",
	HOST, PORT, USER, PASSWORD, DBNAME)

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func (d Database) OpenConnection() *sql.DB {
	db, err = sql.Open(POSTGRESDRIVER, datasourceName)
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Connected!")
	}
	return db
}

func (d Database) CloseConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		return
	}
	log.Println("Disconnected!")
}
