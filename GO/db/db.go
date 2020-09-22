package db

import "database/sql"

func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "mysql123"
	dbName := "classicmodels"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(172.17.0.2:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
