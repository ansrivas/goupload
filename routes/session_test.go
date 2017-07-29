package routes

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SessionTestSuite struct {
	suite.Suite
	store *Store
	req   *http.Request
	res   *httptest.ResponseRecorder
}

func (suite *SessionTestSuite) SetupTest() {

	suite.store, _ = NewDefaultStore(10, "tcp", "localhost:6379", "", []byte("secret-key"))
	suite.req, _ = http.NewRequest("GET", "http://localhost:8080/", nil)
	suite.res = httptest.NewRecorder()
}

func (suite *SessionTestSuite) TearDownTest() {
	suite.store.Close()
	log.Println("Terminating the connection after test to cleanup resoures.")
}

func (suite *SessionTestSuite) Test_GetSession() {

	session, err := suite.store.GetSession(suite.req)

	suite.Nil(err, "Should not get error while getting default session")
	suite.NotNil(session, "Should get sessions successfully")
	suite.Equal(defaultSessionName, session.Name(), "Session names must match")
}

func (suite *SessionTestSuite) Test_WrongParams() {

	store, err := NewDefaultStore(10, "tcp", "localhost:679", "", []byte("secret-key"))

	suite.NotNil(err, "Should get error while getting Store from wrong params")
	suite.Nil(store, "Should fail to sessions successfully")

}

func TestSessionTestSuite(t *testing.T) {
	suite.Run(t, new(SessionTestSuite))
}
