package main

import (
    "github.com/gorilla/mux"
	// "net/http"
    "github.com/codegangsta/negroni"
    "dev/handlers/v1"
    "dev/handlers/v2"
)

func Server() *negroni.Negroni {
	router := mux.NewRouter()
	router.Path("/")
	v1r := router.Headers("Accept", "application/vnd+json").Subrouter()
	v1r.HandleFunc("/json.json", handlersv1.JsonHandler)

	v2r := router.Headers("Accept", "application/vnd.ctemplin.v2+json").Subrouter()
	v2r.HandleFunc("/json.json", handlersv2.JsonHandler)
	
	n := negroni.New()
	n.UseHandler(router)
	return n
}

func main() {
	n := Server()
	n.Run(":9000")
}
