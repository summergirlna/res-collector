//go:build mage
// +build mage

package main

import (
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var (
	Default = Run
	BinPath = filepath.Join("bin", "res-collector")
	DBPath  = filepath.Join("data", "example.db")
	logger  = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
)

// A build step that requires additional params, or platform specific steps for example
func Build() {
	logger.Info("ビルドします")
	cmd := exec.Command("go", "build", "-o", BinPath, ".")
	cmd.Run()
	logger.Info("ビルドしました")
}

// Clean up after yourself
func Clean() {
	logger.Info("バイナリを削除します")
	os.RemoveAll(BinPath)
	logger.Info("バイナリを削除しました")
}

func Run() {
	mg.Deps(Clean)
	mg.Deps(Build)
	cmd := exec.Command(BinPath, "run")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	logger.Info("実行します")
	cmd.Run()
	logger.Info("実行しました")
	mg.Deps(DBSelect)
}

func DBSelect() {
	cmd := exec.Command("sqlite3", DBPath, "SELECT * FROM cpu_stats ORDER BY year DESC, month DESC, day DESC, hour DESC, minute DESC, second DESC LIMIT 1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	logger.Info("ダンプします")
	cmd.Run()
	logger.Info("ダンプしました")
}
