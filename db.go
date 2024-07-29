package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DbConnection() error {
	var DB *sql.DB
	var err error
	conn := "root:root@tcp(localhost:3306)/cetec"
	DB, err = sql.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("database connected")
	return nil
}
