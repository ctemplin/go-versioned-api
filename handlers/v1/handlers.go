package handlersv1

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

type APIv1 struct {
}

func (api *APIv1) Version() string {
	return "v1.0"
}

func (api *APIv1) root_handler(w http.ResponseWriter, r *http.Request) {
	
    fmt.Fprintf(w, "Hi there, I really do love %s!", mux.Vars(r)["path"])
}

func (api *APIv1) JsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"hi": "hello"}` + "\n"))
}

func (api *APIv1) JsonHandler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"hi": "greetings"}` + "\n"))
}

func (api *APIv1) JsonHandler3(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(406)
	w.Write([]byte(`{"error": "This version of the API does not implement this endpoint."}` + "\n"))
}

