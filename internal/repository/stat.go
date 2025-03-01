package repository

import (
	"context"
	"res-collector/internal/model"
)

type Stat interface {
	Save(ctx context.Context, cpu *model.Cpu) error
}
