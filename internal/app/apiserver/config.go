package apiserver

import "github.com/Irishery/golang_webserv.git/internal/app/store"

// Config ...
type Config struct {
	BindAddr string `env:"BIND_ADDR,notEmpty"`
	LogLevel string `env:"LOG_LEVEL,notEmpty"`
	Store    *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
