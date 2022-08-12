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

var Db *sql.DB
var err error
var datasourceName = fmt.Sprintf("host=%s "+
	"port=%s "+
	"user=%s "+
	"password=%s "+
	"dbname=%s "+
	"sslmode=disable",
	HOST, PORT, USER, PASSWORD, DBNAME)

func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func OpenConnection() *sql.DB {
	Db, err = sql.Open(POSTGRESDRIVER, datasourceName)
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Connected!")
	}

	return Db
}

func CloseConnection(db *sql.DB) {
	db.Close()
	log.Println("Disconnected!")
}
