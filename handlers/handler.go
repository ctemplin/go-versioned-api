package api

import "net/http"

type API interface {
	JsonHandler(w http.ResponseWriter, r *http.Request)
	JsonHandler2(w http.ResponseWriter, r *http.Request)
}
