package handlersv2

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "dev/handlers/v1"
)

type APIv2 struct {
	handlersv1.APIv1
}

func (api *APIv2) root_handler(w http.ResponseWriter, r *http.Request) {
	
    fmt.Fprintf(w, "Hi there, I really do love %s!", mux.Vars(r)["path"])
}

func (api *APIv2) JsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"version": 2, "hi": "hello"}` + "\n"))
}

// Commented out as redundant. When not defined the "parent" handler
// is called automatically.
// func (api *APIv2) JsonHandler2(w http.ResponseWriter, r *http.Request) {
// 	api.APIv1.JsonHandler2(w, r)
// }

