package config

import (
    "github.com/joho/godotenv"
    "log"
	"os"
	"fmt"
)

func LoadConfig() {
  
	err := godotenv.Load()
	
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}
	fmt.Println("Environment loaded! ")

}