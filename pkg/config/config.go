package config

import(
	"github.com/akshithere/rssaggregator/internal/db"
	"database/sql"
)

type apiConfig struct {
	Database *sql.DB
}

var ApiConfig apiConfig

func InitConfig(){
	ApiConfig = apiConfig{
	Database: db.DB,
	}
}