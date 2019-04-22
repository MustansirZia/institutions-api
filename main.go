package main

import (
	"fmt"
	"log"
	"os"

	"github.com/qazimusab/musalleen-apis/countries"

	"github.com/qazimusab/musalleen-apis/institutions"
	statesProvider "github.com/qazimusab/musalleen-apis/states/provider"
	"github.com/valyala/fasthttp"
)

var r institutions.InstitutionRepository

func main() {
	// fmt.Println(provider.NewJSONCountryProvider().Provide())
	fmt.Println(statesProvider.NewJSONStateProvider().Provide())

	mux := newMux()
	addr := ":" + getPort()

	fmt.Printf("About to listen on %s!\n", addr)
	if err := fasthttp.ListenAndServe(addr, mux); err != nil {
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
			institutions.ServeHTTP(ctx)
		case "/countries":
			countries.ServeHTTP(ctx)
		default:
			ctx.Error("NOT FOUND", fasthttp.StatusNotFound)
		}
	}
}
