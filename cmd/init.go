package main

import (
	"database/sql"
	"log"
	"stori/database"
)

var pgclient *sql.DB

// initialize database connection
func init() {
	pgclient = database.Open()
	log.Print("postgres successfully initialized")
}
