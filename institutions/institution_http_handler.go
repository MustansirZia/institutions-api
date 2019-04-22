package institutions

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/valyala/fasthttp"
)

var r InstitutionRepository

func init() {
	r = NewInstitutionRepository()
}

func ServeHTTP(ctx *fasthttp.RequestCtx) {
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
	institutions, err := r.GetInstitutions(string(name), count)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		ctx.Error("Oops! Something bad happened.", fasthttp.StatusInternalServerError)
		return
	}
	ctx.Response.Header.SetContentType("application/json")
	if err := json.NewEncoder(ctx).Encode(institutions); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		ctx.Error("Oops! Something bad happened.", fasthttp.StatusInternalServerError)
	}
}
