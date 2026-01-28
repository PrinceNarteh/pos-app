// Package config provides configuration management for the application.
package config

type Config struct {
	App  appConfig
	DB   dbConfig
	Jwt  jwtConfig
	File fileConfig
}

type appConfig struct {
	Port    string
	Env     string
	Version string
}

type dbConfig struct {
	URI          string
	MaxOpenConns int32
	MinOpenConns int32
	MaxIdleTime  string
}

type fileConfig struct {
	UploadPath   string
	MaxSize      int
	ErrorMessage string
	AllowedFiles string
}

type jwtConfig struct {
	Secret           string
	ExpiresIn        string
	RefreshSecret    string
	RefreshExpiresIn string
}
