package main

import (
	// "fmt"
	"encoding/json"
	"testing"
	"net/http"
	"net/http/httptest"
	"net/url"
)


// TestAPIVersionByAcceptHeader checks that values in the request's
// accept header are mapping to the proper version of the API.
func TestAPIVersionByAcceptHeader(t *testing.T) {

	n := Server()

	// Define API versions and their corresponding Accept headers.
	versions := []struct {
		vnum int
		acceptHeader string
	}{
		{1, "application/vnd+json"},
		{2, "application/vnd.ctemplin.v2+json"},
	}

	testUrls := []string{"/json.json", "/json2.json"}

	// For each version make a request and check the version in
	// the response
	for _, version := range versions {

		headers := map[string][]string{
			"Accept": {version.acceptHeader},
		} 

		for _, testUrl := range testUrls {

			response := httptest.NewRecorder()

			url1 := url.URL{Host: "localhost", Path: testUrl}

			request := http.Request{
				URL: &url1,
				Header: headers,
			}

			respObj := struct {
				Version int
				Hi string
			}{}

			n.ServeHTTP(response, &request)
			// fmt.Print(response.Body)
			err := json.Unmarshal(response.Body.Bytes(), &respObj)
			if err != nil {
				t.Error(err)
			}
			if respObj.Version != version.vnum {
				t.Errorf("Wrong API version returned using Accept-Header: %s. Expected: %d, Got: %d", version.acceptHeader, version.vnum, respObj.Version)
			}
		}
	}
}
