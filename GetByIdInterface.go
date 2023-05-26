package persistence

import "context"

type GetByIdInterface interface {
	GetById(ctx context.Context, id string, object interface{}) error
}
