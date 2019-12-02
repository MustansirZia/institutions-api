package api

import (
	"net/http"

	"github.com/mustansirzia/institutions-api/institutions"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	institutions.HandleHTTP(w, r)
}
