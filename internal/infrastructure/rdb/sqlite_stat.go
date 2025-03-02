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

func (S SQLiteStat) SaveCpu(ctx context.Context, cpu *model.Cpu, mng *model.Management) error {
	tx, ok := GetTx(ctx)
	if !ok {
		tx = S.client.MustBegin()
	}
	mp := mapper.NewCpu().From(cpu, mng)

	sql := `
INSERT INTO cpu_stats 
    (stats_id, user_time, system_time, iowait_time) 
VALUES (:stats_id, :user_time, :system_time, :iowait_time)
`
	_, err := tx.NamedExec(sql, &mp)
	if err != nil {
		return err
	}

	return nil
}

func (S SQLiteStat) SaveMemory(ctx context.Context, memory *model.Memory, mng *model.Management) error {
	tx, ok := GetTx(ctx)
	if !ok {
		tx = S.client.MustBegin()
	}
	mp := mapper.NewMemory().From(memory, mng)

	sql := `
INSERT INTO memory_stats 
    (stats_id, application_used) 
VALUES (:stats_id, :application_used)
`
	_, err := tx.NamedExec(sql, &mp)
	if err != nil {
		return err
	}

	return nil
}

func (S SQLiteStat) SaveManagement(ctx context.Context, management *model.Management) error {
	tx, ok := GetTx(ctx)
	if !ok {
		tx = S.client.MustBegin()
	}
	mp := mapper.NewManagement().From(management)

	sql := `
INSERT INTO stats_management 
    (stats_id, year, month, day, hour, minute, second) 
VALUES (:stats_id, :year, :month, :day, :hour, :minute, :second)
`
	_, err := tx.NamedExec(sql, &mp)
	if err != nil {
		return err
	}

	return nil
}
