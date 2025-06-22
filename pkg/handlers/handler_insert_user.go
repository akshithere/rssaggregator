package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"log"
	"github.com/akshithere/rssaggregator/internal/services"
	"github.com/google/uuid"
)

func HandlerInsertUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct{
		Name string `json:"name"`
	}
	var decoder *json.Decoder = json.NewDecoder(r.Body)
	params := parameters{}
	decoder.Decode(&params)
	id := uuid.New()
	res, err := services.InsertUser(w,id,params.Name,time.Now(),time.Now())
	if err != nil{
		log.Fatal("Could not insert a user: ", err)
	}

	log.Println("Inserted a user, result is: ", res)
}
