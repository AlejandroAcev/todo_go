package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	log.Println("Loaded .env file")
	return nil
}

func GetMongoURI() (string, error) {
	err := LoadEnv()

	if err != nil {
		log.Fatalln("Cannot load the MongoURI")
		return "", err
	}

	return os.Getenv("MONGO_URI"), nil
}
