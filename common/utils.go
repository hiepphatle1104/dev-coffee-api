package common

import (
	itemmodel "dev-coffee-api/modules/items/model"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
	"strings"
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

func ValidateImage(image *itemmodel.ItemImage) error {
	if image.URL == "" {
		return errors.New("image url is required")
	}

	parsedURL, err := url.ParseRequestURI(image.URL)
	if err != nil {
		return fmt.Errorf("image.url is not a valid URL")
	}

	if image.Alt == "" {
		return errors.New("image alt is required")
	}

	if !strings.HasPrefix(parsedURL.Scheme, "http") {
		return fmt.Errorf("image.url must start with http or https")
	}

	if image.Width < 100 || image.Width > 1000 {
		return errors.New("width must be greater than 100 and less than 1000")
	}
	if image.Height < 100 || image.Height > 1000 {
		return errors.New("height must be greater than 100 and less than 1000")
	}

	return nil
}
