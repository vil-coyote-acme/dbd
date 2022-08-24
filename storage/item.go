package storage

import "github.com/vil-coyote-acme/dbd/model"

type ItemRepository interface {
	// design considerations :
	// 1 a write is considered success if the write is localy successful (fast return).
	// 2 the cache mechanism is updated asynchrously (but should be fast to read)
	// the replication mechanism should be async
	Set(items ...*model.Item) *Error
}
