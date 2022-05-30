package store

// Config ...
type Config struct {
	POSTGRES_USER     string `env:"DB_USER,notEmpty"`
	POSTGRES_PASSWORD string `env:"DB_PASSWORD,notEmpty"`
	POSTGRES_DB       string `env:"DB_NAME,notEmpty"`
	DATABASE_HOST     string `env:"DB_HOST,notEmpty"`
	DATABASE_PORT     string `env:"DB_PORT,notEmpty"`
	SSL_MODE          string `env:"SSL_MODE,notEmpty"`
}

func NewConfig() *Config {
	return &Config{}
}
