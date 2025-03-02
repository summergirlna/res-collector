package repository

import (
	"context"
	"res-collector/internal/model"
)

type Stat interface {
	SaveCpu(ctx context.Context, cpu *model.Cpu, mng *model.Management) error
	SaveMemory(ctx context.Context, memory *model.Memory, mng *model.Management) error
	SaveManagement(ctx context.Context, management *model.Management) error
}
