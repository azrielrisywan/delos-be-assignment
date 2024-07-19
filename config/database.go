package config

import (
	"log"
	"github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

func SetupDatabase() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=azriel password=Tasik123 dbname=azriel_be_assignment host=postgresql-azriel.alwaysdata.net sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	 // Test the connection to the database
	if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
	 
	return db
}