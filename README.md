# golang_webserv
# .env
BIND_ADDR=:8080  
LOG_LEVEL=debug  

DB_USER=user  
DB_PASSWORD=pas  
DB_HOST=host    
DB_NAME=name  
DB_PORT=5432  
SSL_MODE=disable  

TEST_DB_USER=user  
TEST_DB_PASSWORD=pas  
TEST_DB_HOST=host    
TEST_DB_NAME=name  
TEST_SSL_MODE=disable  
TEST_DB_PORT=5432  

# start up without docker
`make build`  
`make; ./apiserver`
