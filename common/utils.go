package common

import (
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
