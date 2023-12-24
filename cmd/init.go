package main

import (
	"database/sql"
	"log"
	"stori/database"
)

var pgclient *sql.DB

func init() {
	pgclient = database.Open()
	log.Print("postgres successfully initialized")
}
