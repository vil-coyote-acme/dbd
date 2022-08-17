package items

import (
	"encoding/json"
	"net/http"

	"github.com/vil-coyote-acme/dbd/model"
	"github.com/vil-coyote-acme/dbd/security"
	"github.com/vil-coyote-acme/dbd/storage"
)

type Handler struct {
	authenticator  security.HttpAuthenticator
	itemRepository storage.ItemRepository
}

func NewHandler(authenticator security.HttpAuthenticator, itemRepository storage.ItemRepository) *Handler {
	return &Handler{authenticator, itemRepository}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: on POST, do ensure the item payload is complete
	_, err := h.authenticator.LookupSession(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var items []*model.Item
	err = json.NewDecoder(r.Body).Decode(&items)
	if err != nil {
		// TODO send body ?
		// TODO, use 400 instead
		w.WriteHeader(http.StatusInternalServerError)
	}

	h.itemRepository.Set(items...)

	// TODO, returned the items ?
}
