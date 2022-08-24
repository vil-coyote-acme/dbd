package items_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vil-coyote-acme/dbd/model"
	"github.com/vil-coyote-acme/dbd/server/items"
	"github.com/vil-coyote-acme/dbd/utils"
)

func TestNotAuthenticated(t *testing.T) {
	itemsBody := []*model.Item{{Key: "some-key", Data: "some data string"}}
	values, _ := json.Marshal(itemsBody)

	req, err := http.NewRequest("POST", "/items", bytes.NewBuffer(values))
	if err != nil {
		t.Fatalf("could not create a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	h := items.NewSetHandler(utils.AuthenticatorMock{ExpectedError: errors.New("some error")}, &utils.ItemRepositoryMock{})

	// when
	h.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Fatalf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
}

func TestSuccess(t *testing.T) {
	itemsBody := []*model.Item{{Key: "some-key", Data: "some data string"}}

	values, _ := json.Marshal(itemsBody)
	req, err := http.NewRequest("POST", "/items", bytes.NewBuffer(values))
	if err != nil {
		t.Fatalf("could not create a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	itemRepository := &utils.ItemRepositoryMock{}
	h := items.NewSetHandler(utils.AuthenticatorMock{}, itemRepository)

	// when
	h.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if len(itemRepository.GivenPutItems) != 1 {
		t.Errorf("Expect to save %d items, but got %d", 1, len(itemRepository.GivenPutItems))
	}

	if itemsBody[0].Key != itemRepository.GivenPutItems[0].Key {
		t.Errorf("Expect to save item with key %s, but got %s", itemsBody[0].Key, itemRepository.GivenPutItems[0].Key)
	}

	if itemsBody[0].Data != itemRepository.GivenPutItems[0].Data {
		t.Errorf("Expect to save item with data %s, but got %s", itemsBody[0].Data, itemRepository.GivenPutItems[0].Data)
	}
}
