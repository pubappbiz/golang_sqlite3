package main

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
	"strconv"
)

var (
	id int
	languageName string
	development string
)

func dbCRUD() {
	// Open / Close
	db, err := sql.Open("sqlite3", "/Users/pubappbiz/SQLiteData/golang_database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Send ping
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Begin transaction / Rollback
	tran, err :=db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tran.Commit()

	// Create
	sqlCreate, err := db.Prepare("insert into programing_language (language_name, development, created_at, updated_at) values (?, ?, ?, ?);")
	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}
	response, err := sqlCreate.Exec("Go", "Google", time.Now(), time.Now())
	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}
	insertId, _ := response.LastInsertId()
	effected, _ := response.RowsAffected()

	fmt.Println("INSERT id: " + strconv.FormatInt(insertId, 10) + " rows: " + strconv.FormatInt(effected, 10))

	// Read
	sqlRead, err := db.Prepare("select id, language_name, development from programing_language")
	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}
	rows, err := sqlRead.Query()
	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &languageName, &development)
		if err != nil {
			tran.Rollback()
			log.Fatal(err)
		}
		fmt.Println("SELECT id: " + strconv.Itoa(id) + " name: " + languageName + " develop: " + development)
	}

	// Update
	sqlUpdate, err := db.Prepare("update programing_language set language_name = ?, development = ? where id = ?")
	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}
	response, err = sqlUpdate.Exec("PHP", "The PHP Group", insertId)
	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}
	effected, _ = response.RowsAffected()
	fmt.Println("UPDATE rows: " + strconv.FormatInt(effected, 10))

	// Delete
	sqlDelete, err := db.Prepare("delete from programing_language where id = ?")
	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}
	response, err = sqlDelete.Exec(insertId)
	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}
	effected, _ = response.RowsAffected()
	fmt.Println("DELETE rows: " + strconv.FormatInt(effected, 10))

}

func main() {
	dbCRUD()
	fmt.Println("SUCCESS!")
}
