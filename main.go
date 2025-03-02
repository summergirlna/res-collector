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
    stats_id    TEXT PRIMARY KEY,
    user_time   REAL NOT NULL,
    system_time REAL NOT NULL,
    iowait_time REAL NOT NULL,
    FOREIGN KEY (stats_id) REFERENCES stats_management (stats_id)
);

CREATE TABLE IF NOT EXISTS memory_stats
(
    stats_id         TEXT PRIMARY KEY,
    application_used REAL NOT NULL,
    FOREIGN KEY (stats_id) REFERENCES stats_management (stats_id)

);

CREATE TABLE IF NOT EXISTS stats_management
(
    stats_id TEXT PRIMARY KEY,
    year     INTEGER NOT NULL,
    month    INTEGER NOT NULL,
    day      INTEGER NOT NULL,
    hour     INTEGER NOT NULL,
    minute   INTEGER NOT NULL,
    second   INTEGER NOT NULL
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
	memoryGetter := local.NewMemoryGetter()
	service := save_stats.NewProvider(cpuGetter, memoryGetter, stat, transaction)
	if _, err := service.Save(context.TODO(), &save_stats.Input{}); err != nil {
		slog.Error("error occurred", "cause", err.Error())
	}
}
