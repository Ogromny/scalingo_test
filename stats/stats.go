package stats

import (
	"encoding/json"
	"net/http"

	"github.com/Scalingo/go-handlers"
	datatypes "github.com/Scalingo/sclng-backend-test-v1/data_types"
	"github.com/Scalingo/sclng-backend-test-v1/utils"
)

const PATH_PREFIX = "/stats"

// TODO: mutexify ?
var data []datatypes.ReposResponseBodyItems

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

    println(data)

    if err := json.NewEncoder(w).Encode(data); err != nil {
        utils.RespondInternalServerError(err, "Cannot write response body", w, r)
    }
}

func Register(router *handlers.Router) {
	route := router.PathPrefix(PATH_PREFIX)
	route.Methods(http.MethodGet).Path("").HandlerFunc(getHandler)
}

func Consume(items []datatypes.ReposResponseBodyItems) {
    var allowed []datatypes.ReposResponseBodyItems

    for _, item := range items {
        found := false
        for _, _item := range data {
            if _item.Url == item.Url {
                found = true
                break
            }
        }

        if !found {
            allowed = append(allowed, item)
        }
    }

    if allowed != nil {
        data = append(data, allowed...)
    }
}
