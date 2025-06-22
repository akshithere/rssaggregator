package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // _ because solely for side effects
)

var DB *sql.DB
var err error
func ConnectToDB(){
	log.Println("Connecting to the database")
	connectionString := os.Getenv("DB")
	DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal("could not connect to the database", err)
	}
	

	if err = DB.Ping(); err != nil {
		log.Fatal("could not ping the DB ", err)
	} 
}