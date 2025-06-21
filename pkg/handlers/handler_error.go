package handlers

import (
	"net/http"
	"github.com/akshithere/rssaggregator/pkg/helper"
)

func HandlerError(w http.ResponseWriter, r *http.Request){
	helper.RespondWithError(w,500,"something went wrong, interval server error")
}