// Package healthcheck is the package of the healthcheck API route
package healthcheck

import (
	"encoding/json"
	"net/http"
)

// Response its a healthcheck route default response
type Response struct {
	Message string
}

// Handler to healthcheck route
func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	json.NewEncoder(w).Encode(&Response{"Turdus API it's running"})
}
