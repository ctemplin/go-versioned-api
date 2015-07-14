package main

import (
    "github.com/gorilla/mux"
    "github.com/codegangsta/negroni"
    "dev/handlers/v1"
    "dev/handlers/v2"
    "net/http"
    "dev/handlers"
)

func Server() *negroni.Negroni {
	router := mux.NewRouter()
	router.Path("/")

	type version struct {
		acceptHeader string
		apiVersion api.API
	}

	type urlRoute struct {
		path string
		handler func(http.ResponseWriter, *http.Request)
	}
	
	versions := []version {
		{
			"application/vnd+json",
			new(handlersv1.APIv1),
		},
		{
			"application/vnd.ctemplin.v2+json",
			new(handlersv2.APIv2),
		},
	}

	for _, version := range versions {
		subrouter := router.Headers("Accept", version.acceptHeader).Subrouter()
		paths := []string{"/json.json", "/json2.json"}

		pathHandlers := []urlRoute{
			{paths[0], version.apiVersion.JsonHandler},
			{paths[1], version.apiVersion.JsonHandler2},
		}
		for i := 0; i < len(paths); i++ {
			subrouter.HandleFunc(
				pathHandlers[i].path,
				pathHandlers[i].handler)
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
