package security

import "net/http"

type HttpAuthenticator interface {
	LookupSession(r *http.Request) (Session, error)
}

type Session interface {
}
