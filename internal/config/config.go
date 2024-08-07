package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MastodonInstanceURL string
	AccessToken         string
	QuotesFile          string
	SleepDuration       int // in minutes
}

// LoadConfig loads the configuration from the environment variables.
func LoadConfig() (*Config, error) {
	// load from .env file
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}
	intSleepDuration, err := strconv.Atoi(getEnv("SLEEP_DURATION", "300"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse SLEEP_DURATION: %w", err)
	}
	cfg := &Config{
		MastodonInstanceURL: getEnv("MASTODON_INSTANCE_URL", ""),
		AccessToken:         getEnv("ACCESS_TOKEN", ""),
		QuotesFile:          getEnv("QUOTES_FILE", ""),
		SleepDuration:       intSleepDuration,
	}

	if cfg.MastodonInstanceURL == "" {
		return nil, fmt.Errorf("MASTODON_INSTANCE_URL is required")
	}

	if cfg.AccessToken == "" {
		return nil, fmt.Errorf("ACCESS_TOKEN is required")
	}

	if cfg.QuotesFile == "" {
		return nil, fmt.Errorf("POSTS_DIRECTORY is required")
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	value := defaultValue
	if v := os.Getenv(key); v != "" {
		value = v
	}
	return value
}
