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

	type url struct {
		path string
		handler func(http.ResponseWriter, *http.Request)
	}

	type version struct {
		acceptHeader string
		urls []url
	}

	versions := []version {
		{"application/vnd+json",
			[]url{
				{"/json.json", handlersv1.JsonHandler},
				{"/json2.json", handlersv1.JsonHandler2},
			},
		},
		{"application/vnd.ctemplin.v2+json",
			[]url{
				{"/json.json", handlersv2.JsonHandler},
				{"/json2.json", handlersv2.JsonHandler2},
			},
		},
	}

	for _, version := range versions {
		subrouter := router.Headers("Accept", version.acceptHeader).Subrouter()
		for _, url := range version.urls {
			subrouter.HandleFunc(url.path, url.handler)
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
