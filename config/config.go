package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("error loading .env file: ", err)
	}

	return os.Getenv(key)
}
