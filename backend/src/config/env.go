package config

import (
    "github.com/joho/godotenv"
    "log"
	"fmt"
)

func LoadConfig() {
  
	err := godotenv.Load()
	
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Environment loaded! ")
}