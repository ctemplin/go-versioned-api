package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/codegangsta/negroni"
    "dev/handlers"
    "dev/handlers/v1"
    "dev/handlers/v2"
)

func Server() *negroni.Negroni {
	router := mux.NewRouter()
	router.Path("/")

	acceptVersionMap := map[string]api.API {
		"application/vnd+json": new(handlersv1.APIv1),
		"application/vnd.ctemplin.v2+json": new(handlersv2.APIv2),
	}

	for acceptHeader, apiVersion := range acceptVersionMap {
		subrouter := router.Headers("Accept", acceptHeader).Subrouter()

		pathHandlerMap := map[string]func(http.ResponseWriter, *http.Request) {
			"/json.json": apiVersion.JsonHandler,
			"/json2.json": apiVersion.JsonHandler2,
		}
		for path, handler := range pathHandlerMap {
			subrouter.HandleFunc(path, handler)
		}
	}

	n := negroni.New()
	n.UseHandler(router)
	return n
}

func main() {
	n := Server()
	n.Run(":9000")
}
