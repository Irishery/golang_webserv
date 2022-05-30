# golang_webserv
# .env
BIND_ADDR=:8080
LOG_LEVEL=debug

DB_USER=user
DB_PASSWORD=pas
DB_HOST=postgresdb
DB_NAME=name
DB_PORT=5432
SSL_MODE=disable

TEST_DB_USER=postgres
TEST_DB_PASSWORD=pas
TEST_DB_HOST=postgresdb_test
TEST_DB_NAME=currency_info_test
TEST_SSL_MODE=disable
TEST_DB_PORT=5432

# start up without docker
`make build`
`make; ./apiserver`