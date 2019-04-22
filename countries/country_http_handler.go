package countries

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/valyala/fasthttp"
)

var r CountryRepository

func init() {
	r = NewCountryRepository()
}

func ServeHTTP(ctx *fasthttp.RequestCtx) {
	countries, err := r.GetAllCountries()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		ctx.Error("Oops! Something bad happened.", fasthttp.StatusInternalServerError)
		return
	}
	ctx.Response.Header.SetContentType("application/json")
	if err := json.NewEncoder(ctx).Encode(countries); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		ctx.Error("Oops! Something bad happened.", fasthttp.StatusInternalServerError)
	}
}
