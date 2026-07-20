package logger

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSetupSplitsLevels(t *testing.T) {
	originalLogger := slog.Default()
	t.Cleanup(func() {
		slog.SetDefault(originalLogger)
	})

	rootDir := t.TempDir()
	closeLogs, err := Setup(rootDir)
	if err != nil {
		t.Fatalf("Setup() error = %v", err)
	}

	slog.Debug("debug message")
	slog.Info("info message")
	slog.Warn("warn message")
	slog.Error("error message")
	closeLogs()

	debugLog, err := os.ReadFile(filepath.Join(rootDir, "debug.log"))
	if err != nil {
		t.Fatalf("read debug.log: %v", err)
	}
	for _, message := range []string{"debug message", "info message", "warn message", "error message"} {
		if !strings.Contains(string(debugLog), message) {
			t.Errorf("debug.log does not contain %q", message)
		}
	}

	errorLog, err := os.ReadFile(filepath.Join(rootDir, "error.log"))
	if err != nil {
		t.Fatalf("read error.log: %v", err)
	}
	for _, message := range []string{"warn message", "error message"} {
		if !strings.Contains(string(errorLog), message) {
			t.Errorf("error.log does not contain %q", message)
		}
	}
	for _, message := range []string{"debug message", "info message"} {
		if strings.Contains(string(errorLog), message) {
			t.Errorf("error.log unexpectedly contains %q", message)
		}
	}
}
