package main

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type cfg struct {
	DatabaseURL string
	Port        string
}

var configOnce sync.Once
var config cfg

func getConf() cfg {
	configOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		config = cfg{
			DatabaseURL: os.Getenv("DATABASE_URL"),
			Port:        os.Getenv("PORT"),
		}
	})

	return config
}
