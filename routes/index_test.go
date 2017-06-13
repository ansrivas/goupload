package routes

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/ansrivas/goupload/internal"
	"github.com/ansrivas/goupload/logger"
	"github.com/stretchr/testify/suite"
)

type indexTestSuite struct {
	suite.Suite
	routes Resource
}

func (suite *indexTestSuite) SetupTest() {

	templateList := internal.SetAssetsPath("./testdir")
	log := logger.SetupLogger("testapp")

	// We first create the http.Handler we wish to test
	suite.routes = Resource{
		templateList: templateList,
		logger:       log,
	}
}

func (suite *indexTestSuite) TearDownTest() {
	log.Println("This should have been printed after each test to cleanup resoures.")
}

func (suite *indexTestSuite) Test_Get() {
	// We create an http.Request object to test with. The http.Request is
	// totally customizable in every way that a real-life http request is, so
	// even the most intricate behavior can be tested
	r, _ := http.NewRequest("GET", "/", nil)

	// httptest.Recorder implements the http.ResponseWriter interface, and as
	// such can be passed into ServeHTTP to receive the response. It will act as
	// if all data being given to it is being sent to a real client, when in
	// reality it's being buffered for later observation
	w := httptest.NewRecorder()

	suite.routes.GetIndex(w, r)

	// httptest.Recorder gives a number of fields and methods which can be used
	// to observe the response made to our request. Here we check the response
	// code
	suite.Equal(http.StatusOK, w.Code, "/ should return a 200")
}

func (suite *indexTestSuite) Test_Post() {

	// We first create the http.Handler we wish to test

	r, _ := http.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()
	suite.routes.PostIndex(w, r)

	suite.Equal(http.StatusFound, w.Code, "Post request on / should return http.StatusFound.")

	// ------------------------------------------------------------------------------
	form := url.Values{}
	form.Add("emailuser", "admin@gmail.com")
	form.Add("passwordfield", "admin")

	r, _ = http.NewRequest("POST", "/", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	w = httptest.NewRecorder()

	suite.routes.PostIndex(w, r)
	suite.Equal(http.StatusSeeOther, w.Code, "Post request on / should return http.StatusSeeOther.")

}

func TestIndexTestSuite(t *testing.T) {
	suite.Run(t, new(indexTestSuite))
}
