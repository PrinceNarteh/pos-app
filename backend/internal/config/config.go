// Package config provides configuration management for the application.
package config

import "time"

type Config struct {
	App  appConfig
	DB   dbConfig
	Jwt  jwtConfig
	File fileConfig
}

type appConfig struct {
	Port         string
	Env          string
	Version      string
	LogLevel     string
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
	WriteTimeout time.Duration
}

type dbConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	Name            string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	MaxConnLifetime time.Duration
}

type fileConfig struct {
	UploadPath   string
	MaxSize      int
	ErrorMessage string
	AllowedFiles string
}

type jwtConfig struct {
	AccessSecret          string
	AccessExpirationTime  string
	RefreshSecret         string
	RefreshExpirationTime string
}
