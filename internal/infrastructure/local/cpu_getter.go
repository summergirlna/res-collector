package local

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
	"res-collector/internal/component"
	"res-collector/internal/model"
	"time"
)

type CpuGetter struct{}

func NewCpuGetter() component.CpuGetter {
	return &CpuGetter{}
}

func (g CpuGetter) Get(ctx context.Context) (*model.Cpu, error) {
	// MEMO cpu.Times()の引数をtrueにすると、コア毎の値が取得できる
	// MEMO cpu.Timesは動いていた時間を取得する

	// -----
	// 取得開始
	// -----
	start, err := g.getStat()
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Second)
	now := time.Now()
	end, err := g.getStat()
	if err != nil {
		return nil, err
	}

	// -----
	// 差分計算
	// -----
	userDiff, systemDiff, iowaitDiff, totalDiff := g.calculateDiff(start, end)

	return &model.Cpu{
		Timestamp: now,
		User:      userDiff / totalDiff,
		System:    systemDiff / totalDiff,
		IOWait:    iowaitDiff / totalDiff,
	}, nil
}

func (g CpuGetter) getStat() (cpu.TimesStat, error) {
	stats, err := cpu.Times(false)
	if err != nil {
		return cpu.TimesStat{}, err
	}
	if len(stats) == 0 {
		return cpu.TimesStat{}, fmt.Errorf("no cpu times found")
	}

	return stats[0], nil
}

func (g CpuGetter) calculateDiff(start cpu.TimesStat, end cpu.TimesStat) (float64, float64, float64, float64) {
	userDiff := end.User - start.User
	systemDiff := end.System - start.System
	iowaitDiff := end.Iowait - start.Iowait
	totalDiff := (end.User + end.System + end.Iowait + end.Idle) -
		(start.User + start.System + start.Iowait + start.Idle)

	return userDiff, systemDiff, iowaitDiff, totalDiff
}
