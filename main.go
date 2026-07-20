package main

import (
	"log/slog"
	"os"

	"github.com/novmbrs/passwords/cmd"
	"github.com/novmbrs/passwords/internal/config"
	"github.com/novmbrs/passwords/internal/logger"
)

func main() {
	os.Exit(run())
}

func run() int {
	if err := config.LoadEnv(); err != nil {
		slog.Error("load environment", "error", err)
		return 1
	}

	rootDir, err := config.DefaultRootDir()
	if err != nil {
		slog.Error("find passwords directory", "error", err)
		return 1
	}

	closeLogs, err := logger.Setup(rootDir)
	if err != nil {
		slog.Error("set up logging", "error", err)
		return 1
	}
	defer closeLogs()

	return cmd.Execute()
}
