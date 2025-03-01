package save_stats

import (
	"context"
	"log/slog"
	"res-collector/internal/component"
	"res-collector/internal/repository"
)

type Provider struct {
	cpuGetter component.CpuGetter
	stats     repository.Stat
}

func NewProvider(cpuGetter component.CpuGetter, stats repository.Stat) Service {
	return &Provider{cpuGetter: cpuGetter, stats: stats}
}

func (p Provider) Save(ctx context.Context, in *Input) (*Output, error) {
	slog.Info("start getting cpu stats")
	c, err := p.cpuGetter.Get(ctx)
	if err != nil {
		return nil, err
	}

	slog.Info("end getting cpu stats", "stats", c.String())

	slog.Info("save start")
	err = p.stats.Save(ctx, c)
	if err != nil {
		return nil, err
	}
	slog.Info("save end")

	return &Output{}, nil
}
