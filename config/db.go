package config

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"

	"github.com/jaydeep87/poc-go-pg/controllers"
)

// Connect function
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "password",
		Addr:     "localhost:5432",
		Database: "postgres",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateUserTable(db)
	controllers.InitiateDB(db)
	return db
}
