package save_stats

import (
	"context"
	"log/slog"
	"res-collector/internal/component"
	"res-collector/internal/repository"
)

type Provider struct {
	cpuGetter   component.CpuGetter
	stats       repository.Stat
	transaction repository.Transaction
}

func NewProvider(cpuGetter component.CpuGetter, stats repository.Stat, transaction repository.Transaction) Service {
	return &Provider{cpuGetter: cpuGetter, stats: stats, transaction: transaction}
}

func (p Provider) Save(ctx context.Context, in *Input) (*Output, error) {
	slog.Info("start getting cpu stats")
	c, err := p.cpuGetter.Get(ctx)
	if err != nil {
		return nil, err
	}

	slog.Info("end getting cpu stats", "stats", c.String())

	slog.Info("save start")
	_, err = p.transaction.Do(ctx, func(ctx context.Context) (interface{}, error) {
		err = p.stats.Save(ctx, c)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	slog.Info("save end")

	return &Output{}, nil
}
