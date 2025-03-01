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

func (S SQLiteStat) Save(ctx context.Context, cpu *model.Cpu) error {
	mp := mapper.FromCpu(cpu)

	tx := S.client.MustBegin()
	sql := `
INSERT INTO cpu_stats 
    (serial_number, year, month, day, hour, minute, second, user_time, system_time, iowait_time) 
VALUES (:serial_number, :year, :month, :day, :hour, :minute, :second, :user_time, :system_time, :iowait_time)
`
	_, err := tx.NamedExec(sql, &mp)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
