package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func Open(username, password, database, host, port string) {
	initialize(username, password, database, host, port)
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", username, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func Close() {
	_ = DB.Close()
}

func Migrate() {

}

func initialize(username, password, database, host, port string) {
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/", username, password, host, port))
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query("CREATE DATABASE IF NOT EXISTS " + database)
	if err != nil {
		log.Fatal(err)
	}

	_ = db.Close()
}
