package main

import (
	"log"

	"github.com/Irishery/golang_webserv.git/internal/app/apiserver"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var (
	configPath string
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config := apiserver.NewConfig()
	if err := env.Parse(config); err != nil {
		log.Print("main conf")
		log.Fatalf("%+v", err)
	}

	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
