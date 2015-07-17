package main

import (
	// "fmt"
	"encoding/json"
	"testing"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// Define API versions and their corresponding Accept headers.
var versions = map[string]string {
	"v1.0": "application/vnd+json",
	"v2.0": "application/vnd.example.v2+json",
	"v3.0": "application/vnd.example.v3+json",
}

var testPaths = []string{"/json.json", "/json2.json", "/json3.json"}

// TestAPIVersionByAcceptHeader checks that values in the request's
// accept header are mapping to the proper version of the API.
func TestAPIVersionByAcceptHeader(t *testing.T) {

	n := Server()

	// For each version make a request and check the version in
	// the response
	for versionString, acceptHeader := range versions {

		headers := map[string][]string{
			"Accept": {acceptHeader},
		} 

		for _, testPath := range testPaths {

			response := httptest.NewRecorder()

			url1 := url.URL{Host: "localhost", Path: testPath}

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

			responseVersion, exists := response.Header()["X-Example-Version"]
			if !exists {
				t.Error("API version missing from response headers.")
			}
			if versionString != responseVersion[0]  {
				t.Errorf("Wrong API version returned using Accept-Header: %s. Expected: %d, Got: %d", acceptHeader, versionString, responseVersion[0])
			}
		}
	}
}

// TestAPIVersionByAcceptHeader checks that values in the request's
// accept header are mapping to the proper version of the API.
func TestAPIVersionByQueryString(t *testing.T) {

	n := Server()

	// For each version make a request and check the version in
	// the response
	for versionString, acceptHeader := range versions {

		headers := map[string][]string{
			"Accept": {"*/*"},
		} 

		for _, testPath := range testPaths {

			response := httptest.NewRecorder()

			url1 := url.URL{Host: "localhost", Path: testPath, RawQuery: "apiv="+versionString}

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

			responseVersion, exists := response.Header()["X-Example-Version"]
			if !exists {
				t.Error("API version missing from response headers.")
			}
			if versionString != responseVersion[0]  {
				t.Errorf("Wrong API version returned using Accept-Header: %s. Expected: %d, Got: %d", acceptHeader, versionString, responseVersion[0])
			}
		}
	}
}


// TestNotAcceptable checks that the 406 response is returned when proper.
func TestNotAcceptable(t *testing.T) {

	n := Server()

	response := httptest.NewRecorder()

	headers := map[string][]string{
		"Accept": {versions["v1.0"]},
	}  

	url1 := url.URL{Host: "localhost", Path: "/json3.json"}

	request := http.Request{
		URL: &url1,
		Header: headers,
	}

	n.ServeHTTP(response, &request)
	// fmt.Print(response.Body)
	if response.Code != 406 {
		t.Errorf("Wrong response code returned using Accept-Header: %s. Expected: %d, Got: %d", versions["v1.0"], 406, response.Code)

	}
}
