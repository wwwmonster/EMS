package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEvnVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .evn file")
	}
}
