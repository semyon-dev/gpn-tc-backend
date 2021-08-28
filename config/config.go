package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	MongoUrl        string
	Port            string
	ParseHabrCareer string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("can't load from file: " + err.Error())
	}
	MongoUrl = os.Getenv("MONGO_URL")
	Port = os.Getenv("PORT")
	ParseHabrCareer = os.Getenv("PARSE_HABR_CAREER")
}
