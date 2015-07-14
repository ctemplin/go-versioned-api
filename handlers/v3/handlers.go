package handlersv3

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "dev/handlers/v1"
)

type APIv3 struct {
	handlersv1.APIv1
}

func (api *APIv3) Version() string {
	return "v3.0"
}

func (api *APIv3) root_handler(w http.ResponseWriter, r *http.Request) {
	
    fmt.Fprintf(w, "Hi there, I really do love %s!", mux.Vars(r)["path"])
}

func (api *APIv3) JsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"hi": "hello"}` + "\n"))
}

func (api *APIv3) JsonHandler3(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"hi": "salutations"}` + "\n"))
}
