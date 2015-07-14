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

	for acceptHeader, vApi := range acceptVersionMap {
		// Create a subrouter for the header/api version.
		subrouter := router.Headers("Accept", acceptHeader).Subrouter()

		// Define the path/handler relationships.
		pathHandlerMap := map[string]func(http.ResponseWriter, *http.Request) {
			"/json.json": vApi.JsonHandler,
			"/json2.json": vApi.JsonHandler2,
		}
		// Create a route in the subrouter for each path/handler.
		for path, handler := range pathHandlerMap {
			subrouter.HandleFunc(path, handler)
		}
	}

	n := negroni.New()
	n.UseHandlerFunc(ApiVersionMiddleware)
	n.UseHandler(router)
	return n
}

var acceptVersionMap = map[string]api.API {
	"application/vnd+json": new(handlersv1.APIv1),
	"application/vnd.ctemplin.v2+json": new(handlersv2.APIv2),
}

func ApiVersionMiddleware(w http.ResponseWriter, r *http.Request) {
	accept := r.Header["Accept"]
	api, exists := acceptVersionMap[accept[0]]
	var version string
	if exists {
		version = api.Version()
	} else {
		version = "unknown"
	}
	w.Header().Set("X-ctemplin-version", version)
}

func main() {
	n := Server()
	n.Run(":9000")
}
