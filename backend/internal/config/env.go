package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

func initConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("could not load .env file")
	}

	return &Config{
		App: appConfig{
			Env:          getEnv("APP_ENV", "development"),
			Version:      getEnv("APP_VERSION", "1.0.0"),
			Port:         fmt.Sprintf(":%s", getEnv("APP_PORT", "3000")),
			ReadTimeout:  getEnvAsDuration("READ_TIMEOUT", 10*time.Second),
			WriteTimeout: getEnvAsDuration("WRITE_TIMEOUT", 10*time.Second),
			IdleTimeout:  getEnvAsDuration("IDLE_TIMEOUT", 60*time.Second),
			LogLevel:     getEnv("LOG_LEVEL", "info"),
		},
		DB: dbConfig{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnv("DB_PORT", "5432"),
			User:            getEnv("DB_USER", "postgres"),
			Password:        getEnv("DB_PASSWORD", "postgres"),
			Name:            getEnv("DB_NAME", "pos-app"),
			SSLMode:         getEnv("DB_SSL_MODE", "disable"),
			MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 10),
			MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 100),
			MaxConnLifetime: getEnvAsDuration("DB_MAX_CONN_LIFETIME", time.Hour),
		},
		Jwt: jwtConfig{
			AccessSecret:          getEnv("JWT_SECRET", ""),
			AccessExpirationTime:  getEnv("JWT_EXPIRES_IN", "100s"),
			RefreshSecret:         getEnv("JWT_REFRESH_SECRET", ""),
			RefreshExpirationTime: getEnv("JWT_REFRESH_EXPIRES_IN", "3600s"),
		},
		File: fileConfig{
			UploadPath:   getEnv("FILE_UPLOAD_PATH", ""),
			MaxSize:      getEnvAsInt("FILE_MAX_SIZE", 5242880),
			ErrorMessage: getEnv("FILE_MAX_MESSAGE", "File size must be less that 5MB"),
			AllowedFiles: getEnv("FILE_ALLOWED_TYPES", ""),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	valueStr := getEnv(key, "")

	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return fallback
}

func getEnvAsDuration(key string, fallback time.Duration) time.Duration {
	valueStr := getEnv(key, "")
	if value, err := time.ParseDuration(valueStr); err == nil {
		return value
	}
	return fallback
}
