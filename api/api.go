// Package api is the package of the Turdus API interface
package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cantoniazzi/turdus/api/healthcheck"
	"github.com/gorilla/mux"
)

// Start is the function responsible to start and setup the Turdus API
func Start() {
	port := os.Getenv("API_PORT")
	r := mux.NewRouter()

	r.HandleFunc("/health-check",
		healthcheck.Handler,
	).Methods("GET")

	fmt.Println("Turdus API it's running on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
