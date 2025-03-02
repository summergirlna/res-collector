package repository

import (
	"context"
	"res-collector/internal/model"
)

type Stat interface {
	SaveCpu(ctx context.Context, cpu *model.Cpu) error
	SaveMemory(ctx context.Context, memory *model.Memory) error
}
