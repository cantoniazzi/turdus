// Package tweet_test is responsible to test tweet route
package tweet_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cantoniazzi/turdus/api/tweet"
	"github.com/cantoniazzi/turdus/service"
	"github.com/jarcoal/httpmock"

	"github.com/gorilla/mux"
)

func TestHandler(t *testing.T) {
	url := "url.fake"
	os.Setenv("TWITTER_API_URL", url) // TODO
	os.Setenv("TWITTER_API_TOKEN_BEARER", "asasas")
	httpmock.Activate()

	r := mux.NewRouter()
	w := httptest.NewRecorder()

	payload := `{"author": "superman","dateFrom": "202001010001","dateTo": "202001101159"}`

	// TODO: see a way to create a factory to mock external api tests
	tr := service.TweetResponse{
		Results: []service.Tweet{
			{CreatedAt: "123", ID: 102, Text: "A nice tweet", FavoriteCount: 20, RetweetCount: 10},
		},
	}
	trEncoded, _ := json.Marshal(tr)
	httpmock.RegisterResponder(
		"POST",
		url,
		httpmock.NewStringResponder(200, string(trEncoded)),
	)

	r.HandleFunc("/tweets", tweet.Handler)
	r.ServeHTTP(w, httptest.NewRequest(
		"POST",
		"/tweets",
		bytes.NewBufferString(payload),
	))

	// TODO: improve this test checking the route reponse as well
	if w.Code != http.StatusOK {
		t.Error("Unexpected http status code", w.Code)
	}
}
