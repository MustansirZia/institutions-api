package institutions

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/qazimusab/musalleen-apis/institutions/providers"

	"github.com/valyala/fasthttp"
)

var r InstitutionRepository
var r2 InstitutionRepository

func init() {
	r = NewInstitutionRepository(
		providers.NewIndianCollegesProvider(),
		providers.NewWorldUniversitiesProvider(),
		providers.NewIndianUniversitiesProvider(),
	)
}

func ServeHTTP(ctx *fasthttp.RequestCtx) {
	name := ctx.QueryArgs().Peek("name")
	if len(name) == 0 {
		ctx.Error("query param \"name\" is missing. Please try again.", fasthttp.StatusBadRequest)
		return
	}
	t := time.Now().Unix()
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
	fmt.Println(time.Now().Unix() - t)
	ctx.Response.Header.SetContentType("application/json")
	if err := json.NewEncoder(ctx).Encode(institutions); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		ctx.Error("Oops! Something bad happened.", fasthttp.StatusInternalServerError)
	}
}
