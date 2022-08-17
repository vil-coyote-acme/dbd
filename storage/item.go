package storage

import "github.com/vil-coyote-acme/dbd/model"

type ItemRepository interface {
	Set(items ...*model.Item) *Error
}
