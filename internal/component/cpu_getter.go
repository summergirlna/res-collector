package component

import (
	"context"
	"res-collector/internal/model"
)

type CpuGetter interface {
	Get(ctx context.Context) (*model.Cpu, error)
}
