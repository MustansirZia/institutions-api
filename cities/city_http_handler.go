package cities

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/valyala/fasthttp"
)

var r CityRepository

func init() {
	r = NewCityRepository()
}

func ServeHTTPByName(ctx *fasthttp.RequestCtx) {
	nameBytes := ctx.QueryArgs().Peek("name")

	if len(nameBytes) == 0 {
		ctx.Error("query param \"name\" must be present. Please try again.", fasthttp.StatusBadRequest)
		return
	}

	cities, err := r.GetCitiesByName(string(nameBytes))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		ctx.Error("Oops! Something bad happened.", fasthttp.StatusInternalServerError)
		return
	}
	ctx.Response.Header.SetContentType("application/json")
	if err := json.NewEncoder(ctx).Encode(cities); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		ctx.Error("Oops! Something bad happened.", fasthttp.StatusInternalServerError)
	}
}

func ServeHTTPByState(ctx *fasthttp.RequestCtx) {
	nameBytes := ctx.QueryArgs().Peek("state")

	if len(nameBytes) == 0 {
		ctx.Error("query param \"state\" must be present. Please try again.", fasthttp.StatusBadRequest)
		return
	}

	cities, err := r.GetCitiesByState(string(nameBytes))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		ctx.Error("Oops! Something bad happened.", fasthttp.StatusInternalServerError)
		return
	}
	ctx.Response.Header.SetContentType("application/json")
	if err := json.NewEncoder(ctx).Encode(cities); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		ctx.Error("Oops! Something bad happened.", fasthttp.StatusInternalServerError)
	}
}
