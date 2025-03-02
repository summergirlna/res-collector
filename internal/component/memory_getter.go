package component

import (
	"context"
	"res-collector/internal/model"
)

type MemoryGetter interface {
	Get(ctx context.Context) (*model.Memory, error)
}
