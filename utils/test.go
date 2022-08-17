package utils

import (
	"net/http"
	"testing"

	"github.com/vil-coyote-acme/dbd/model"
	"github.com/vil-coyote-acme/dbd/security"
	"github.com/vil-coyote-acme/dbd/storage"
)

type AuthenticatorMock struct {
	ExpectedSession security.Session
	ExpectedError   error
}

func (a AuthenticatorMock) LookupSession(r *http.Request) (security.Session, error) {
	return a.ExpectedSession, a.ExpectedError
}

type ItemRepositoryMock struct {
	GivenPutItems    []*model.Item
	ExpectedPutError *storage.Error
	t                *testing.T
}

func (m *ItemRepositoryMock) Set(items ...*model.Item) *storage.Error {
	m.GivenPutItems = append(m.GivenPutItems, items...)
	return m.ExpectedPutError
}
