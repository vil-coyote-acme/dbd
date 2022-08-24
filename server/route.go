package server

import (
	"github.com/julienschmidt/httprouter"
	"github.com/vil-coyote-acme/dbd/server/items"
)

func InitRouter() *httprouter.Router {

	router := httprouter.New()
	// TODO add required implementations
	router.Handler("POST", "/api/v1/items", items.NewSetHandler(nil, nil))

	return router
}
