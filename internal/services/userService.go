package services

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/akshithere/rssaggregator/pkg/config"
	"github.com/akshithere/rssaggregator/pkg/helper"
	"github.com/google/uuid"
)

func CreateUserTable() {
	log.Printf("Creating User Table")

	query := `CREATE TABLE IF NOT EXISTS USERS(
	UUID UUID PRIMARY KEY,
	NAME TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL)`

	config.ApiConfig.Database.Exec(query)

}

func InsertUser(w http.ResponseWriter, uuid uuid.UUID, name string, created_at time.Time, updated_at time.Time) (sql.Result, error) {
	query := `INSERT INTO USERS(UUID, NAME, created_at, updated_at)
	VALUES($1, $2, $3, $4)`
	res, err := config.ApiConfig.Database.Exec(query, uuid, name, created_at, updated_at)
	if err != nil {
		helper.RespondWithError(w, 500, "Could not write to the database")
		return nil, err
	}
	fmt.Println("The result of the exec command is:", res)
	return res, nil
}
