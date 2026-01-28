package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Env = initConfig()

func initConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("could not load .env file")
	}

	return &Config{
		App: appConfig{
			Port:    fmt.Sprintf(":%s", getEnv("APP_PORT", "3000")),
			Env:     getEnv("APP_ENV", "development"),
			Version: getEnv("APP_VERSION", "1.0.0"),
		},
		DB: dbConfig{
			URI:          getEnv("DB_URI", ""),
			MaxOpenConns: int32(getEnvAsInt("DB_MAX_OPEN_CONNS", 10)),
			MinOpenConns: int32(getEnvAsInt("DB_MAX_IDLE_CONNS", 10)),
			MaxIdleTime:  getEnv("MAX_IDLE_TIME", "10m"),
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
	valueStr := getEnv(key, "0")

	if valueStr == "" {
		return fallback
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return fallback
	}
	return value
}
