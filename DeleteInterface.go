package persistence

import "context"

type DeleteInterface interface {
	Delete(ctx context.Context, id string) error
}
