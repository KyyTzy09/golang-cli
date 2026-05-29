package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	GeminiAPIKey string `mapstructure:"gemini_api_key"`
	GroqAPIKey   string `mapstructure:"groq_api_key"`
}

var AppConfig Config

func LoadConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Path config
	configDir := filepath.Join(home, ".config", "ai-cli")

	// Buat folder kalau belum ada
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return err
	}

	viper.AddConfigPath(configDir)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.SetDefault("default_model", "gemini")

	if err := viper.ReadInConfig(); err != nil {
		// Kalau filenya belum ada, buat file kosong baru
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := viper.WriteConfigAs(filepath.Join(configDir, "config.json")); err != nil {
				return fmt.Errorf("gagal membuat file config baru: %w", err)
			}
		} else {
			return fmt.Errorf("gagal membaca config: %w", err)
		}
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("gagal parsing config: %w", err)
	}

	return nil
}
