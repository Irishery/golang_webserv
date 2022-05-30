package store_test

type TestingConfig struct {
	POSTGRES_USER     string `env:"TEST_DB_USER,notEmpty"`
	POSTGRES_PASSWORD string `env:"TEST_DB_PASSWORD,notEmpty"`
	POSTGRES_DB       string `env:"TEST_DB_NAME,notEmpty"`
	DATABASE_HOST     string `env:"TEST_DB_HOST,notEmpty"`
	DATABASE_PORT     string `env:"TEST_DB_PORT,notEmpty"`
	SSL_MODE          string `env:"TEST_SSL_MODE,notEmpty"`
}

func NewTestConfig() *TestingConfig {
	return &TestingConfig{}
}
