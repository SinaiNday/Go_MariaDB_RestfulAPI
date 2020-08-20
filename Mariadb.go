package main

import (
	
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB){
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "sinai"
    dbName := "restfulapi"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
	return db
}
