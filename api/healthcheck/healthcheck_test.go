package healthcheck_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cantoniazzi/turdus/api/healthcheck"

	"github.com/gorilla/mux"
)

func TestHandler(t *testing.T) {
	r := mux.NewRouter()
	w := httptest.NewRecorder()

	r.HandleFunc("/health-check", healthcheck.Handler)
	r.ServeHTTP(w, httptest.NewRequest("GET", "/health-check", nil))

	if w.Code != http.StatusOK {
		t.Error("Unexpected http status code", w.Code)
	}

}
