package service_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/cantoniazzi/turdus/service"
	"github.com/jarcoal/httpmock"
)

func getServiceMock() service.TwitterService {
	return service.TwitterService{
		URL:         "www.twitter_api.com",
		TokenBearer: "a secret bearer token",
		MaxResults:  100,
	}
}

func TestGetTweets(t *testing.T) {
	httpmock.Activate()

	svcMock := getServiceMock()

	var author = "author"
	var fromDate = "202001010001"
	var toDate = "202001101159"
	var nextPage = ""

	tr := service.TweetResponse{
		Next: "",
		Results: []service.Tweet{
			{CreatedAt: "123", ID: 102, Text: "A nice tweet", FavoriteCount: 20, RetweetCount: 10},
		},
	}

	// TODO: see a way to create a factory to mock external api tests
	trEncoded, _ := json.Marshal(tr)
	httpmock.RegisterResponder(
		"POST",
		svcMock.URL,
		httpmock.NewStringResponder(200, string(trEncoded)),
	)

	tweets, _ := svcMock.GetTweets(author, fromDate, toDate, nextPage)
	if reflect.DeepEqual(tweets.Results, &tr.Results) {
		t.Error("Unexpected get tweets response", tweets)
	}
}
