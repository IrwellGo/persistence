package persistence

import "context"

type CreateInterface interface {
	Create(ctx context.Context, id string, object interface{}) error
}
