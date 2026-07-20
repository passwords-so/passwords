package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

// Setup configures the default slog logger and returns a function that closes its files.
func Setup(rootDir string) (func(), error) {
	if err := os.MkdirAll(rootDir, 0o700); err != nil {
		return nil, fmt.Errorf("create log directory: %w", err)
	}

	debugFile, err := openLogFile(filepath.Join(rootDir, "debug.log"))
	if err != nil {
		return nil, err
	}

	errorFile, err := openLogFile(filepath.Join(rootDir, "error.log"))
	if err != nil {
		_ = debugFile.Close()
		return nil, err
	}

	debugHandler := slog.NewTextHandler(debugFile, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	errorHandler := slog.NewTextHandler(errorFile, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	})

	slog.SetDefault(slog.New(slog.NewMultiHandler(debugHandler, errorHandler)))

	return func() {
		_ = errorFile.Close()
		_ = debugFile.Close()
	}, nil
}

func openLogFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return nil, fmt.Errorf("open log file %q: %w", path, err)
	}
	if err := file.Chmod(0o600); err != nil {
		_ = file.Close()
		return nil, fmt.Errorf("secure log file %q: %w", path, err)
	}
	return file, nil
}
