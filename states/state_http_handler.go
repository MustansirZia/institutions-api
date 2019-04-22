package states

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/valyala/fasthttp"
)

var r StateRepository

func init() {
	r = NewStateRepository()
}

func ServeHTTP(ctx *fasthttp.RequestCtx) {
	country := ctx.QueryArgs().Peek("country")
	if len(country) == 0 {
		ctx.Error("query param \"country\" is missing. Please try again.", fasthttp.StatusBadRequest)
		return
	}
	states, err := r.GetStates(string(country))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		ctx.Error("Oops! Something bad happened.", fasthttp.StatusInternalServerError)
		return
	}
	ctx.Response.Header.SetContentType("application/json")
	if err := json.NewEncoder(ctx).Encode(states); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		ctx.Error("Oops! Something bad happened.", fasthttp.StatusInternalServerError)
	}
}
