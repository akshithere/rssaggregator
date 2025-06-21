package services

import ( 
	"log"
	"github.com/akshithere/rssaggregator/pkg/config"
)

func CreateUserTable(){
	log.Printf("Creating User Table")

	query := `CREATE TABLE IF NOT EXISTS USERS(
	UUID SERIAL PRIMARY KEY,
	NAME TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL)`

	config.ApiConfig.Database.Exec(query)

}