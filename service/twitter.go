// Package service is the service layer of the turdus project
package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const maxResultsPerPage int = 100

// ExtendedTweet represents a extended tweet unitary structure
type ExtendedTweet struct {
	FullText string `json:"full_text"`
}

// Tweet represents a tweet unitary structure
type Tweet struct {
	CreatedAt     string        `json:"created_at"`
	ID            int           `json:"id"`
	Text          string        `json:"text"`
	FavoriteCount int           `json:"favorite_count"`
	RetweetCount  int           `json:"retweet_count"`
	ExtendedTweet ExtendedTweet `json:"extended_tweet"`
}

// TweetPayload is the payload to sent to twitter api
type TweetPayload struct {
	Query      string `json:"query"`
	MaxResults int    `json:"maxResults"`
	FromDate   string `json:"fromDate"`
	ToDate     string `json:"toDate"`
	Next       string `json:"next,omitempty"` // Informs if has next page
}

// TweetResponse is the twitter api response
type TweetResponse struct {
	Results []Tweet `json:"results"`
	Next    string  `json:"next,omitempty"` // Informs if has next page
}

// TwitterService represents the twitter service structure
type TwitterService struct {
	URL         string
	TokenBearer string
	MaxResults  int
}

// NewTwitterService is a TwitterService constructor
func NewTwitterService() TwitterService {
	return TwitterService{
		URL:         os.Getenv("TWITTER_API_URL"),
		TokenBearer: os.Getenv("TWITTER_TOKEN_BEARER"),
		MaxResults:  maxResultsPerPage,
	}
}

// GetTweets retrieve tweets from the twitter api given specific search
func (svc TwitterService) GetTweets(author, fromDate, toDate, nextPage string) (*TweetResponse, error) {
	var authorParam = fmt.Sprintf("from:%s lang:pt", author)
	var output TweetResponse

	payload := TweetPayload{
		Query:      authorParam,
		MaxResults: svc.MaxResults,
		FromDate:   fromDate,
		ToDate:     toDate,
	}

	if nextPage != "" {
		payload.Next = nextPage
	}

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error when trying to parse the payload. %v", err)
		return nil, err
	}

	client := &http.Client{}
	r, err := http.NewRequest(http.MethodPost, svc.URL, bytes.NewBuffer(data))
	r.Header.Add("Authorization", svc.TokenBearer)
	if err != nil {
		log.Printf("Error when trying to prepare the post request. %v", err)
		return nil, err
	}

	res, err := client.Do(r)
	if err != nil || res.StatusCode != 200 {
		if err == nil {
			err = fmt.Errorf(res.Status)
		}

		log.Printf("Error when trying to execute the post request. %v", err)
		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(&output); err != nil {
		log.Printf("Error when trying to parse the api response body. %v", err)
		return nil, err
	}

	defer res.Body.Close()

	return &output, nil
}
