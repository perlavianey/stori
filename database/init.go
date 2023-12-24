package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
)

// Open Database
func Open() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
		"db", "5432", "postgres", "p0stgr3s.D4t4b4s3", "postgres")
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}
	db.SetMaxOpenConns(20)
	return db
}

func StartDB() (*pg.DB, error) {
	var (
		opts *pg.Options
		err  error
	)

	opts = &pg.Options{
		Addr:     "composepostgres:5432",
		User:     "postgres",
		Password: "p0stgr3s.D4t4b4s3",
		Database: "postgres",
	}

	//connect db
	db := pg.Connect(opts)
	//run migrations
	collection := migrations.NewCollection()
	err = collection.DiscoverSQLMigrations("migrations")
	if err != nil {
		return nil, err
	}

	//start the migrations
	_, _, err = collection.Run(db, "init")
	if err != nil {
		return nil, err
	}

	oldVersion, newVersion, err := collection.Run(db, "up")
	if err != nil {
		return nil, err
	}
	if newVersion != oldVersion {
		log.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		log.Printf("version is %d\n", oldVersion)
	}

	//return the db connection
	return db, err
}
