package api

import "net/http"

type API interface {
	Version() string
	JsonHandler(w http.ResponseWriter, r *http.Request)
	JsonHandler2(w http.ResponseWriter, r *http.Request)
	JsonHandler3(w http.ResponseWriter, r *http.Request)
}
