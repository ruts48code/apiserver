package main

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	utils "github.com/ruts48code/utils4ruts"
)

func getDBS() (*sql.DB, error) {
	dbN := utils.RandomArrayString(conf.DBS)
	dbConnect := false
	var db *sql.DB
	var err error
	qstring := ""
	for i := range dbN {
		switch conf.DBType {
		case "postgres":
			if conf.DBPassword == "" {
				qstring = "postgres://" + conf.DBUsername + "@" + dbN[i] + "/" + conf.DBName
			} else {
				qstring = "postgres://" + conf.DBUsername + ":" + conf.DBPassword + "@" + dbN[i] + "/" + conf.DBName
			}
			if conf.DBParam != "" {
				qstring = qstring + "?" + conf.DBParam
			}
		case "mysql":
			if conf.DBPassword == "" {
				qstring = conf.DBUsername + "@tcp(" + dbN[i] + ":3306)/" + conf.DBName
			} else {
				qstring = conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + dbN[i] + ":3306)/" + conf.DBName
			}
			if conf.DBParam != "" {
				qstring = qstring + "?" + conf.DBParam
			}
		}

		db, err = sql.Open(conf.DBType, qstring)
		if err != nil {
			log.Printf("Error: Fail to open db %s - %v\n", dbN[i], err)
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Printf("Error: Fail to ping db %s - %v\n", dbN[i], err)
			db.Close()
			continue
		}
		log.Printf("Log: Connect to db %s\n", dbN[i])
		dbConnect = true
		break
	}
	if !dbConnect {
		log.Printf("Error: Cannot connect to all db\n")
		return nil, errors.New("cannot connect to all db")
	}
	return db, nil
}
