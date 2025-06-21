package handlers

import (
	"net/http"
	"github.com/akshithere/rssaggregator/pkg/helper"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request){
	helper.RespondWithJSON(w, struct{}{}, 200)
}