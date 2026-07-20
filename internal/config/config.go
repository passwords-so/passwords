package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	VaultsDir string `mapstructure:"vaults_dir"`
}

var C Config

// Load configures Viper and reads an optional config file.
func Load() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("find user config directory: %w", err)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(filepath.Join(configDir, "passwords"))
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) {
			return fmt.Errorf("read config: %w", err)
		}
	}

	if err := viper.Unmarshal(&C); err != nil {
		return fmt.Errorf("decode config: %w", err)
	}

	return nil
}
