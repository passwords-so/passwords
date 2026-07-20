package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

type Config struct {
	VaultsDir string `mapstructure:"vaults_dir"`
}

var C Config
var IsDev bool

// Load configures Viper and reads an optional config file.
func Load() error {
	if err := setDefaults(); err != nil {
		return err
	}

	if err := ensureConfigFile(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&C); err != nil {
		return fmt.Errorf("decode config: %w", err)
	}

	if err := os.MkdirAll(C.VaultsDir, 0o700); err != nil {
		return fmt.Errorf("create vaults directory: %w", err)
	}

	return nil
}

func LoadEnv() error {
	if err := gotenv.Load(); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("load .env: %w", err)
	}

	if env := os.Getenv("PASSWORDS_ENV"); env != "" {
		IsDev = env == "dev"
	}
	return nil
}

func ensureConfigFile() error {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("find user config directory: %w", err)
	}

	configDir := filepath.Join(userConfigDir, "passwords")
	configPath := filepath.Join(configDir, "config.yaml")

	if err := os.MkdirAll(configDir, 0o700); err != nil {
		return fmt.Errorf("create config directory: %w", err)
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	viper.SetConfigPermissions(0o600)

	if err := viper.ReadInConfig(); err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("read config: %w", err)
		}
		if err := viper.SafeWriteConfigAs(configPath); err != nil {
			return fmt.Errorf("create config: %w", err)
		}
	}

	return nil
}

func setDefaults() error {
	dataDir, err := defaultDataDir()
	if err != nil {
		return err
	}

	viper.SetDefault(
		"vaults_dir",
		filepath.Join(dataDir, "passwords", "vaults"),
	)

	return nil
}

func defaultDataDir() (string, error) {
	if dataDir := os.Getenv("XDG_DATA_HOME"); dataDir != "" {
		return dataDir, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("find user home directory: %w", err)
	}

	return filepath.Join(homeDir, ".local", "share"), nil
}
