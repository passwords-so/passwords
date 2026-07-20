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
	RootDir   string `mapstructure:"-"`
	VaultsDir string `mapstructure:"vaults_dir"`
}

var C Config
var IsDev bool

// Load configures Viper and reads an optional config file.
func Load() error {
	rootDir, err := DefaultRootDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(rootDir, 0o700); err != nil {
		return fmt.Errorf("create passwords directory: %w", err)
	}

	configPath := filepath.Join(rootDir, "config.yaml")
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	viper.SetConfigPermissions(0o600)
	viper.SetDefault("vaults_dir", filepath.Join(rootDir, "vaults"))

	if err := ensureConfigFile(configPath); err != nil {
		return err
	}

	if err := viper.Unmarshal(&C); err != nil {
		return fmt.Errorf("decode config: %w", err)
	}
	C.RootDir = rootDir

	if err := os.MkdirAll(C.VaultsDir, 0o700); err != nil {
		return fmt.Errorf("create vaults directory: %w", err)
	}

	return nil
}

func LoadEnv() error {
	if err := gotenv.Load(); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("load .env: %w", err)
	}

	IsDev = os.Getenv("PASSWORDS_ENV") == "dev"
	return nil
}

func ensureConfigFile(configPath string) error {
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

// DefaultRootDir returns the directory containing all passwords application files.
func DefaultRootDir() (string, error) {
	var (
		baseDir string
		err     error
	)

	if IsDev {
		baseDir, err = os.Getwd()
	} else {
		baseDir, err = os.UserHomeDir()
	}
	if err != nil {
		return "", fmt.Errorf("find passwords root directory: %w", err)
	}

	return filepath.Join(baseDir, ".passwords"), nil
}
