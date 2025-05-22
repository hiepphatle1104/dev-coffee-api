package common

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvLookup(key string) string {
	data, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalln("Missing environment variable: " + key)
	}

	return data
}

func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file")
	}
}
