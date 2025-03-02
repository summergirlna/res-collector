package rdb

import (
	"context"
	"github.com/jmoiron/sqlx"
	"res-collector/internal/infrastructure/rdb/mapper"
	"res-collector/internal/model"
	"res-collector/internal/repository"
)

type SQLiteStat struct {
	client *sqlx.DB
}

func NewSQLiteStat(client *sqlx.DB) repository.Stat {
	return &SQLiteStat{
		client: client,
	}
}

func (S SQLiteStat) SaveCpu(ctx context.Context, cpu *model.Cpu) error {
	tx, ok := GetTx(ctx)
	if !ok {
		tx = S.client.MustBegin()
	}
	mp := mapper.NewCpu().From(cpu)

	sql := `
INSERT INTO cpu_stats 
    (serial_number, year, month, day, hour, minute, second, user_time, system_time, iowait_time) 
VALUES (:serial_number, :year, :month, :day, :hour, :minute, :second, :user_time, :system_time, :iowait_time)
`
	_, err := tx.NamedExec(sql, &mp)
	if err != nil {
		return err
	}

	return nil
}

func (S SQLiteStat) SaveMemory(ctx context.Context, memory *model.Memory) error {
	tx, ok := GetTx(ctx)
	if !ok {
		tx = S.client.MustBegin()
	}
	mp := mapper.NewMemory().From(memory)

	sql := `
INSERT INTO memory_stats 
    (serial_number, year, month, day, hour, minute, second, application_used) 
VALUES (:serial_number, :year, :month, :day, :hour, :minute, :second, :application_used)
`
	_, err := tx.NamedExec(sql, &mp)
	if err != nil {
		return err
	}

	return nil
}
