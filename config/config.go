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
	ParseSuppliers  string
	ParseRbk        string
	ParseOkved      string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("can't load .env file")
	}
	MongoUrl = os.Getenv("MONGO_URL")
	Port = os.Getenv("PORT")
	ParseHabrCareer = os.Getenv("PARSE_HABR_CAREER")
	ParseSuppliers = os.Getenv("PARSE_SUPPLIERS")
	ParseRbk = os.Getenv("PARSE_RBK")
	ParseOkved = os.Getenv("PARSE_OKVED")
}
