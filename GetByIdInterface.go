package persistence

import "context"

type GetByIdInterface interface {
	GetById(ctx context.Context, id string, toObject *interface{}) error
}
