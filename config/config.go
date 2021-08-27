package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	MongoUrl string
	Port     string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("can't load from file: " + err.Error())
	}
	MongoUrl = os.Getenv("MONGO_URL")
	Port = os.Getenv("PORT")
}
