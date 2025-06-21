package handlers

import (
	"net/http"
	"github.com/akshithere/rssaggregator/internal/services"
)

func HandlerCreateUserTable(w http.ResponseWriter, r *http.Request){
	services.CreateUserTable()
}