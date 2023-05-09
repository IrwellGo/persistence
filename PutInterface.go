package persistence

import "context"

type PutInterface interface {
	Put(ctx context.Context, id string, object interface{}) error
}
