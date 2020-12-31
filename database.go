package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pass123"
	dbname   = "notes_dev"
)

func connectDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}

func migrateDb(db *sql.DB) {
	sqlStatement := `
	CREATE TABLE notes (
		id SERIAL PRIMARY KEY,
		body TEXT
	  );
	  `

	_, err := db.Exec(sqlStatement)

	if err != nil {
		// panic(err)
		fmt.Println(err)
	}
}
