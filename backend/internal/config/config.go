// Package config provides configuration management for the application.
package config

type Config struct {
	APP appConfig
	DB  dbConfig
}

type appConfig struct {
	Port string
	Env  string
}

type dbConfig struct {
	URI          string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}
