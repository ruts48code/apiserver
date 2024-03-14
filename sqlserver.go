package main

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func getSQLServer() (db *sql.DB, err error) {
	db, err = sql.Open("sqlserver", conf.Personal.Server)
	if err != nil {
		log.Printf("Error: Can not connect to SQL server\n")
	} else {
		log.Printf("Log: Connect to SQL server\n")
	}
	return
}
