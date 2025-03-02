package repository

import "context"

type Transaction interface {
	Do(ctx context.Context, f func(ctx context.Context) (any, error)) (any, error)
}
