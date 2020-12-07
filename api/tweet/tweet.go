// Package tweet is related with get tweets route
package tweet

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cantoniazzi/turdus/service"
)

// Response represents the tweet route response
type Response struct {
	Results []service.Tweet `json:"tweets"`
}

// Request represents the tweets route default request
type Request struct {
	Author   string `json:"author"`
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`
}

// Handler to tweets route
func Handler(w http.ResponseWriter, r *http.Request) {
	// TODO: validate/format request fields (date params)
	body := Request{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("Error when trying to parse tweet request. %v", err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	svc := service.NewTwitterService()
	tweets, err := svc.GetTweets(
		body.Author,
		body.DateFrom,
		body.DateTo,
		"",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer r.Body.Close()

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(&Response{tweets.Results})
}
