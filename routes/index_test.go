package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Get(t *testing.T) {

	// We first create the http.Handler we wish to test
	route := IndexResource{}.GetIndex

	// We create an http.Request object to test with. The http.Request is
	// totally customizable in every way that a real-life http request is, so
	// even the most intricate behavior can be tested
	r, _ := http.NewRequest("GET", "/", nil)

	// httptest.Recorder implements the http.ResponseWriter interface, and as
	// such can be passed into ServeHTTP to receive the response. It will act as
	// if all data being given to it is being sent to a real client, when in
	// reality it's being buffered for later observation
	w := httptest.NewRecorder()

	route(w, r)

	// httptest.Recorder gives a number of fields and methods which can be used
	// to observe the response made to our request. Here we check the response
	// code
	if w.Code != 200 {
		t.Fatalf("wrong code returned: %d", w.Code)
	}

	// We can also get the full body out of the httptest.Recorder, and check
	// that its contents are what we expect
	body := w.Body.String()
	assert.Equal(t, "welcome", body, "Get request on / should return welcome.")

}

func Test_Post(t *testing.T) {

	// We first create the http.Handler we wish to test
	route := IndexResource{}.PostIndex

	r, _ := http.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()
	route(w, r)

	if w.Code != 200 {
		t.Fatalf("wrong code returned: %d", w.Code)
	}

	body := w.Body.String()
	assert.Equal(t, "welcome", body, "Get request on / should return welcome.")

}
