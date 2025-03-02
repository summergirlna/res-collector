package local

import (
	"context"
	"github.com/shirou/gopsutil/v4/mem"
	"res-collector/internal/component"
	"res-collector/internal/model"
)

type MemoryGetter struct{}

func NewMemoryGetter() component.MemoryGetter {
	return &MemoryGetter{}
}

func (m MemoryGetter) Get(ctx context.Context) (*model.Memory, error) {
	stats, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	available := stats.Available
	total := stats.Total

	return &model.Memory{
		Used: float64(total-available) / float64(total) * percentage,
	}, nil
}
