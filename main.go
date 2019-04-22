package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/qazimusab/musalleen-apis/institutions"
	"github.com/qazimusab/musalleen-apis/institutions/providers"
	statesProvider "github.com/qazimusab/musalleen-apis/states/provider"
	"github.com/valyala/fasthttp"
)

var r institutions.InstitutionRepository

func main() {
	// fmt.Println(provider.NewJSONCountryProvider().Provide())
	fmt.Println(statesProvider.NewJSONStateProvider().Provide())
	loadRepository()

	mux := newMux()
	addr := ":" + getPort()

	fmt.Printf("About to listen on %s!\n", addr)
	if err := fasthttp.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func loadRepository() {
	r = institutions.NewInstitutionRepository(
		// Exisiting providers
		providers.NewIndianCollegesProvider(),
		providers.NewIndianUniversitiesProvider(),
		providers.NewWorldUniversitiesProvider(),
		// Add your own `InstitutionProvider` instance here.
	)
	if err := r.Load(); err != nil {
		log.Fatal(err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	return port
}

func newMux() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/institutions":
			handleInstitutionSearch(ctx)
		default:
			ctx.Error("NOT FOUND", fasthttp.StatusNotFound)
		}
	}
}

func handleInstitutionSearch(ctx *fasthttp.RequestCtx) {
	name := ctx.QueryArgs().Peek("name")
	if len(name) == 0 {
		ctx.Error("query param \"name\" is missing. Please try again.", fasthttp.StatusBadRequest)
		return
	}
	count, err := ctx.QueryArgs().GetUint("count")
	if err != nil {
		// Default count is 10.
		count = 10
	}
	institutions := r.GetInstitutions(string(name), count)
	ctx.Response.Header.SetContentType("application/json")
	if err := json.NewEncoder(ctx).Encode(institutions); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		ctx.Error("Oops! Something bad happened.", fasthttp.StatusInternalServerError)
	}
}
