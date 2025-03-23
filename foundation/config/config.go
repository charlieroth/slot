package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

type WebConfig struct {
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
	Port            int
}

type DBConfig struct {
	URL      string
	MinConns int32
	MaxConns int32
}

type Config struct {
	Web WebConfig
	DB  DBConfig
}

func LoadConfig() (*Config, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, errors.New("DATABASE_URL is not set")
	}

	minConnsStr := os.Getenv("DB_MIN_CONNS")
	if minConnsStr == "" {
		return nil, errors.New("DB_MIN_CONNS is not set")
	}

	minConns, err := strconv.Atoi(minConnsStr)
	if err != nil {
		return nil, errors.New("DB_MIN_CONNS is not a valid integer")
	}

	maxConnsStr := os.Getenv("DB_MAX_CONNS")
	if maxConnsStr == "" {
		return nil, errors.New("DB_MAX_CONNS is not set")
	}

	maxConns, err := strconv.Atoi(maxConnsStr)
	if err != nil {
		return nil, errors.New("DB_MAX_CONNS is not a valid integer")
	}

	return &Config{
		Web: WebConfig{
			ReadTimeout:     10 * time.Second,
			WriteTimeout:    10 * time.Second,
			IdleTimeout:     10 * time.Second,
			ShutdownTimeout: 10 * time.Second,
			Port:            3000,
		},
		DB: DBConfig{
			URL:      dbURL,
			MinConns: int32(minConns),
			MaxConns: int32(maxConns),
		},
	}, nil
}
