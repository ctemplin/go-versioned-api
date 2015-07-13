package handlersv1

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func root_handler(w http.ResponseWriter, r *http.Request) {
	
    fmt.Fprintf(w, "Hi there, I really do love %s!", mux.Vars(r)["path"])
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"version": 1, "hi": "bye"}` + "\n"))
}

