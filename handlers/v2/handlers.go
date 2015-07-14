package handlersv2

import (
    "net/http"
    "dev/handlers/v1"
)

type APIv2 struct {
	handlersv1.APIv1
}

func (api *APIv2) Version() string {
	return "v2.0"
}

func (api *APIv2) JsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"hi": "hello"}` + "\n"))
}

// Commented out as redundant. When not defined the "parent" handler
// is called automatically.
// func (api *APIv2) JsonHandler2(w http.ResponseWriter, r *http.Request) {
// 	api.APIv1.JsonHandler2(w, r)
// }

