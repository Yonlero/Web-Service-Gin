package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	POSTGRESDRIVER = "postgres"
	USER           = "postgres"
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

func OpenConnection() {
	Db, err = sql.Open(POSTGRESDRIVER, datasourceName)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("\n\nConnected!\n\n")
	}
}

func CloseConnection(db *sql.DB) {
	db.Close()
	fmt.Printf("\n\nDatabase closed\n\n")
}
