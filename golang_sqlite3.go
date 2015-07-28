package main

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func dbConnectionAndClose() {
	db, err := sql.Open("sqlite3", "/Users/pubappbiz/SQLiteData/golang_database")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	dbConnectionAndClose()
	fmt.Println("SUCCESS!")
}
