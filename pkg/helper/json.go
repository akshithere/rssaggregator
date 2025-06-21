package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, payload interface{}, code int){
	data , err := json.Marshal(payload)
	if err != nil{
		log.Fatal("something's up from our side", err);
		w.WriteHeader(500);
		return;
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func RespondWithError(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Println("something went wrong", msg)
	}
	type ErrorResponse struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w,ErrorResponse{
		Error: msg,
	},code)
}