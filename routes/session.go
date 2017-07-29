package routes

import (
	"log"
	"net/http"

	"github.com/boj/redistore"
	"github.com/gorilla/sessions"
)

var defaultSessionName = "github.com.ansrivas.goupload"

// Store ..
type Store struct {
	Redis       *redistore.RediStore
	SessionName string
}

// NewDefaultStore creates a Store object.
func NewDefaultStore(poolsize int, network string, address string, password string, keyPairs ...[]byte) (*Store, error) {

	// Fetch new store.
	store, err := redistore.NewRediStore(poolsize, network, address, password, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &Store{Redis: store}, nil
}

// Close closes a redis connection.
func (s *Store) Close() {
	s.Redis.Close()
}

// GetSession returns a gorilla session for an incoming http.Request.
func (s *Store) GetSession(r *http.Request) (*sessions.Session, error) {
	if s.SessionName == "" {
		log.Printf("Session name was not set. Using default name %s", defaultSessionName)
		s.SetSession(defaultSessionName)
	}

	return s.Redis.Get(r, s.SessionName)
}

// SetSession sets a name for a session.
func (s *Store) SetSession(name string) {
	s.SessionName = name
}
