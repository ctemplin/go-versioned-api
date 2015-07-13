package main

import (
    "github.com/gorilla/mux"
    "github.com/codegangsta/negroni"
    "dev/handlers/v1"
    "dev/handlers/v2"
    "net/http"
)

func Server() *negroni.Negroni {
	router := mux.NewRouter()
	router.Path("/")

	versions := []struct {
		handler func(http.ResponseWriter, *http.Request)
		acceptHeader string
	}{
		{handlersv1.JsonHandler, "application/vnd+json"},
		{handlersv2.JsonHandler, "application/vnd.ctemplin.v2+json"},
	}

	for _, version := range versions {
		subrouter := router.Headers("Accept", version.acceptHeader).Subrouter()
		subrouter.HandleFunc("/json.json", version.handler)
	}
	
	n := negroni.New()
	n.UseHandler(router)
	return n
}

func main() {
	n := Server()
	n.Run(":9000")
}
