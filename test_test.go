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
	versions := map[string]string {
		"v1.0": "application/vnd+json",
		"v2.0": "application/vnd.ctemplin.v2+json",
		"v3.0": "application/vnd.ctemplin.v3+json",
	}

	testUrls := []string{"/json.json", "/json2.json", "/json3.json"}

	// For each version make a request and check the version in
	// the response
	for versionString, acceptHeader := range versions {

		headers := map[string][]string{
			"Accept": {acceptHeader},
		} 

		for _, testUrl := range testUrls {

			response := httptest.NewRecorder()

			url1 := url.URL{Host: "localhost", Path: testUrl}

			request := http.Request{
				URL: &url1,
				Header: headers,
			}

			respObj := struct {
				Hi string
			}{}

			n.ServeHTTP(response, &request)
			// fmt.Print(response.Body)
			err := json.Unmarshal(response.Body.Bytes(), &respObj)
			if err != nil {
				t.Error(err)
			}

			responseVersion, exists := response.Header()["X-Ctemplin-Version"]
			if !exists {
				t.Error("API version missing from response headers.")
			}
			if versionString != responseVersion[0]  {
				t.Errorf("Wrong API version returned using Accept-Header: %s. Expected: %d, Got: %d", acceptHeader, versionString, responseVersion[0])
			}
		}
	}
}
