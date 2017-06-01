package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Static(t *testing.T) {

	route := StaticResource{staticDir: "./testdir"}.Routes().ServeHTTP
	r, _ := http.NewRequest("GET", "/test.css", nil)
	w := httptest.NewRecorder()

	route(w, r)
	assert.Equal(t, w.Code, http.StatusOK, "Static file should be served with 200ok ")

	body := w.HeaderMap.Get("Content-Type")
	expected := "text/css; charset=utf-8"
	assert.Equal(t, body, expected, "Get request on / should return css.")

}
