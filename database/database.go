package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

var DB *sql.DB

func Open(username, password, database, host, port string) {
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", username, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func Close() {
	DB.Close()
}