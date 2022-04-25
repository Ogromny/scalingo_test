package repos

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Scalingo/go-handlers"
	"github.com/Scalingo/sclng-backend-test-v1/data_types"
	"github.com/Scalingo/sclng-backend-test-v1/stats"
	"github.com/Scalingo/sclng-backend-test-v1/utils"
)

const (
	API_URL     = "https://api.github.com"
	PATH_PREFIX = "/repos"
)

// Unfortunately cannot be made const ?
var (
	_metakeys = []string{"stars", "language", "user", "org", "size", "forks", "license", "archived"}
	_username string
	_password string
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	// NOTE: Create request
	request, err := http.NewRequest(http.MethodGet, API_URL+"/search/repositories", nil)
	if err != nil {
		utils.RespondInternalServerError(err, "Cannot create request", w, r)
		return
	}
	request.Header.Set("Accept", "application/vnd.github.v3+json")

    if _username != "" && _password != "" {
        request.SetBasicAuth(_username, _password)
    }

	// NOTE: Construct q
	var buffer strings.Builder
	for _, meta := range _metakeys {
		if values, ok := r.URL.Query()[meta]; ok {
			for _, value := range values {
				buffer.WriteString(meta + ":" + value + " ")
			}
		}
	}

	if texts, ok := r.URL.Query()["text"]; ok {
		buffer.WriteString(strings.Join(texts, ""))
	}

	// NOTE: Construct query
	query := request.URL.Query()
	query.Add("per_page", "100")
	query.Add("q", buffer.String())
	if value, ok := r.URL.Query()["page"]; ok {
		query.Add("page", value[0])
	}

	request.URL.RawQuery = query.Encode()

	// NOTE: Make the request
	client := http.Client{Timeout: time.Duration(time.Second * 5)}
	response, err := client.Do(request)
	if err != nil {
		utils.RespondInternalServerError(err, "Cannot make the API request", w, r)
		return
	}

	// NOTE: consume JSON
	defer response.Body.Close()

	var data datatypes.ReposResponseBody
	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		utils.RespondInternalServerError(err, "Cannot decode API response", w, r)
		return
	}

	go stats.Consume(data.Items)

	// NOTE: Write response
	w.Header().Add("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(data); err != nil {
		utils.RespondInternalServerError(err, "Cannot write response", w, r)
	}
}

func Register(router *handlers.Router) {
	route := router.PathPrefix(PATH_PREFIX)
	route.Methods(http.MethodGet).Path("").HandlerFunc(getHandler)
}

func SetToken(username string, password string) {
	_username = username
	_password = password
}
