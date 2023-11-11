package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	postgresDNS = "postgres://"
	driver      = "postgres"
)

func main() {
	_, err := connectToDB(driver)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("test")
}

func connectToDB(driver string) (*sql.DB, error) {
	db, err := sql.Open(driver, postgresDNS)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
