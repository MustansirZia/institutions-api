package institutions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/mustansirzia/institutions-api/institutions/providers"
)

var repository InstitutionRepository
var r2 InstitutionRepository

func init() {
	repository = NewInstitutionRepository(
		providers.NewIndianCollegesProvider(),
		providers.NewWorldUniversitiesProvider(),
		providers.NewIndianUniversitiesProvider(),
	)
}

func HandleHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		writeErr(w, http.StatusBadRequest, "query param \"name\" is missing. Please try again.")
		return
	}
	count := r.URL.Query().Get("count")
	countInt, err := strconv.Atoi(count)
	if err != nil {
		// Default count is 10.
		countInt = 10
	}
	institutions, err := repository.GetInstitutions(string(name), countInt)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		writeErr(w, http.StatusInternalServerError, "Oops! Something bad happened.")
		return
	}
	w.Header().Add("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(institutions); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		writeErr(w, http.StatusInternalServerError, "Oops! Something bad happened.")
	}
}

func writeErr(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprint(w, message)
}
