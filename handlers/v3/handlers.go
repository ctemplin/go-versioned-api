package handlersv3

import (
	"net/http"
	"dev/handlers/v2"
)

type APIv3 struct {
	handlersv2.APIv2
}

func (api *APIv3) Version() string {
	return "v3.0"
}

func (api *APIv3) JsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"greeting": "hello"}` + "\n"))
}

func (api *APIv3) JsonHandler3(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"greeting": "salutations"}` + "\n"))
}
