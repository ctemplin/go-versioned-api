package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"github.com/mistifyio/negroni-pprof"
	"dev/handlers"
	"dev/handlers/v1"
	"dev/handlers/v2"
	"dev/handlers/v3"
	"fmt"
	"os"
)

var acceptVersionMap = map[string]api.API {
	"application/vnd+json": new(handlersv1.APIv1),
	"application/vnd.example.v2+json": new(handlersv2.APIv2),
	"application/vnd.example.v3+json": new(handlersv3.APIv3),
}

var AcceptVersionMap = func() map[string]api.API {
	return acceptVersionMap
}

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router = AddRoutes(router)
  return router
}

func AddRoutes(router *mux.Router) *mux.Router{
	router.Path("/")

	for acceptHeader, vApi := range acceptVersionMap {
		// Create a subrouter for the header/api version.
		subrouter := router.MatcherFunc(
			acceptOrQueryMatcherFactory(acceptHeader)).Subrouter()

		// Define the path/handler relationships.
		pathHandlerMap := map[string]func(http.ResponseWriter, *http.Request) {
			"/json.json": vApi.JsonHandler,
			"/json2.json": vApi.JsonHandler2,
			"/json3.json": vApi.JsonHandler3,
		}
		// Create a route in the subrouter for each path/handler.
		for path, handler := range pathHandlerMap {
			route := subrouter.HandleFunc(path, handler)
			route.Name(fmt.Sprintf("%s - %s", path, handler))
		}
	}
	return router
}


var queryVersionMap = map[string]api.API {
	"1": new(handlersv1.APIv1),
	"2": new(handlersv2.APIv2),
	"3": new(handlersv3.APIv3),
}

func ApiVersionMiddleware(w http.ResponseWriter, r *http.Request) {
	accept := r.Header["Accept"]
	api, exists := acceptVersionMap[accept[0]]
	var version string
	if exists {
		version = api.Version()
	} else {
		version = r.FormValue("apiv")
	}
	w.Header().Set("X-example-version", version)
}

func ContentTypeMiddleware(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func acceptOrQueryMatcherFactory(acceptHeader string) (func(*http.Request, *mux.RouteMatch) bool) {
	return func(r *http.Request, rm *mux.RouteMatch) bool {
		var isHeaderMatch bool = r.Header["Accept"][0] == acceptHeader
		var isQueryStringMatch bool = r.FormValue("apiv") == acceptVersionMap[acceptHeader].Version()
		return isHeaderMatch || isQueryStringMatch
	}
}

func Server() *negroni.Negroni {
	n := negroni.New()
	n.Use(pprof.Pprof())
	n.UseHandlerFunc(ContentTypeMiddleware)
	n.UseHandlerFunc(ApiVersionMiddleware)
	n.UseHandler(CreateRouter())
	return n
}

func main() {
	n := Server()
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	n.Run(fmt.Sprintf(":%s", port))
}
