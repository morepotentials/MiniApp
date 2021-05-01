package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {

	type expected struct {
		status int
		// body   string
	}

	type test struct {
		name     string
		useDB    bool
		url      string
		expected expected
		filepath string
		eval     func([]byte) (bool, string)
	}

	mockTests := []test{
		{
			name:     "route not found",
			url:      "/doesnt/exist",
			filepath: "testdata/users.json",
			expected: expected{
				status: http.StatusNotFound,
			},
		},
		{
			name:     "simple users request foo exists",
			url:      "/users",
			filepath: "testdata/foo.json",
			expected: expected{
				status: http.StatusOK,
			},
			eval: func(body []byte) (bool, string) {
				users := make([]User, 0)
				err := json.Unmarshal(body, &users)
				if err != nil {
					return false, "could not unmarshal body into users"
				}
				for _, user := range users {
					if user.FirstName == "Foo" {
						return true, ""
					}
				}
				return false, "could not find Foo in users body response"
			},
		},
		//TODO: write a new test to call get broker's clients endpoint.
		// pass a custom url with a broker id.
		// expect status 200
		// write custom eval function for the test, assert on existence
		// based on user id. i.e, just user id 234 exists in response.
	}
	dbTests := []test{
		{
			name:  "simple users request foo exists",
			useDB: true,
			url:   "/users",
			expected: expected{
				status: http.StatusOK,
			},
			eval: func(body []byte) (bool, string) {
				users := make([]User, 0)
				err := json.Unmarshal(body, &users)
				if err != nil {
					return false, "could not unmarshal body into users"
				}
				for _, user := range users {
					if user.FirstName == "Erik" {
						return true, ""
					}
				}
				return false, "could not find Erik in users body response"
			},
		},
	}

	tests := make([]test, 0, len(mockTests)+len(dbTests))
	tests = append(tests, mockTests...)
	tests = append(tests, dbTests...)

	for _, test := range tests {
		// for each test, reinitialize fetch with new filepath
		testService := Service{}
		if test.useDB {
			testService.Fetcher = newDatabaseUsersFetcher()
		} else {
			testService.Fetcher = JSONRefresherUsersFetcher{filename: test.filepath}
		}

		testRouter := testService.NewRouter()

		req, err := http.NewRequest("GET", test.url, nil)
		if err != nil {
			t.Fatal(err)
		}
		req.RequestURI = test.url
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()

		t.Run(test.name, func(f *testing.T) {
			// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			testRouter.ServeHTTP(rr, req)
			// Check the status code is what we expect.
			if rr.Code != test.expected.status {
				f.Errorf("handler returned wrong status code: got %v want %v", rr.Code, test.expected.status)
			}

			if rr.Code == http.StatusOK { // only run this for tests with 200 status code..
				pass, cause := test.eval(rr.Body.Bytes())
				if pass != true {
					f.Errorf("eval func did not pass, cause: %s", cause)
				}
			}

		})
	}
}
