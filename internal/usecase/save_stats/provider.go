package save_stats

import (
	"context"
	"log/slog"
	"res-collector/internal/component"
	"res-collector/internal/repository"
)

type Provider struct {
	cpuGetter    component.CpuGetter
	memoryGetter component.MemoryGetter
	stats        repository.Stat
	transaction  repository.Transaction
}

func NewProvider(
	cpuGetter component.CpuGetter,
	memoryGetter component.MemoryGetter,
	stats repository.Stat,
	transaction repository.Transaction,
) Service {
	return &Provider{
		cpuGetter:    cpuGetter,
		memoryGetter: memoryGetter,
		stats:        stats,
		transaction:  transaction,
	}
}

func (p Provider) Save(ctx context.Context, in *Input) (*Output, error) {
	slog.Info("start getting cpu stats")
	c, err := p.cpuGetter.Get(ctx)
	if err != nil {
		return nil, err
	}
	slog.Info("end getting cpu stats", "stats", c.String())

	slog.Info("start getting mem stats")
	m, err := p.memoryGetter.Get(ctx)
	if err != nil {
		return nil, err
	}
	slog.Info("end getting mem stats", "stats", m.String())

	slog.Info("save start")
	_, err = p.transaction.Do(ctx, func(ctx context.Context) (interface{}, error) {
		err = p.stats.SaveCpu(ctx, c)
		if err != nil {
			return nil, err
		}
		err = p.stats.SaveMemory(ctx, m)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	slog.Info("save end")

	return &Output{}, nil
}
