package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"path/filepath"
	"res-collector/internal/infrastructure/local"
	"res-collector/internal/infrastructure/rdb"
	"res-collector/internal/usecase/save_stats"

	_ "github.com/glebarez/go-sqlite"
)

const (
	schema = `
CREATE TABLE IF NOT EXISTS cpu_stats
(
    serial_number TEXT PRIMARY KEY,
    year          INTEGER NOT NULL,
    month         INTEGER NOT NULL,
    day           INTEGER NOT NULL,
    hour          INTEGER NOT NULL,
    minute        INTEGER NOT NULL,
    second        INTEGER NOT NULL,
    user_time     REAL    NOT NULL,
    system_time   REAL    NOT NULL,
    iowait_time   REAL    NOT NULL
);
`
)

func main() {
	db, err := sqlx.Connect("sqlite", filepath.Join("data", "example.db"))
	if err != nil {
		slog.Error("failed to connect to database")
		return
	}
	db.MustExec(schema)

	stat := rdb.NewSQLiteStat(db)
	transaction := rdb.NewTransaction(db)
	cpuGetter := local.NewCpuGetter()
	service := save_stats.NewProvider(cpuGetter, stat, transaction)
	if _, err := service.Save(context.TODO(), &save_stats.Input{}); err != nil {
		slog.Error("error occurred", "cause", err.Error())
	}
}
