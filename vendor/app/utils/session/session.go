package session

import (
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// Store - store for session
	Store *sessions.CookieStore
)

// Session ...
type Session struct {
	SecretKey     string
	EncryptionKey string
	Options       sessions.Options
}

// InitSession ...
func InitSession(s *Session, domain string) {
	Store = sessions.NewCookieStore([]byte(s.SecretKey))
	Store.Options = &s.Options
}

// CheckUserInSession - check if user in session
func CheckUserInSession(s *sessions.Session) bool {
	return s.Values["id"] != nil
}

// GetUserID ...
func GetUserID(s *sessions.Session) (uint, error) {
	id, ok := s.Values["id"].(uint)
	if ok {
		return id, nil
	}

	return 0, errors.New("Not integer value")
}

// Instance returns a new session, never returns an error
func Instance(r *http.Request) *sessions.Session {
	// id := r.Header.Get("X-Request-ID")
	// r.Header.Set("X-Request-ID", id)
	// log.Println(r.Header, id)
	session, _ := Store.Get(r, "session")
	return session
}

// Clear deletes all the current session values
func Clear(s *sessions.Session) {
	for k := range s.Values {
		delete(s.Values, k)
	}
}

// Expire - delete cookie
func Expire(s *sessions.Session) {
	Clear(s)
	s.Options = &sessions.Options{
		MaxAge: -1,
	}
}
