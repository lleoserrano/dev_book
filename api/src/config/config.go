package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringDataBaseConnection = ""
	Port                     = 8080
	SecretKey                []byte
)

// Set the environment variable in the application
func Init() {
	var e error

	if e = godotenv.Load(); e != nil {
		log.Fatal(e)
	}

	Port, e = strconv.Atoi(os.Getenv("API_PORT"))
	if e != nil {
		Port = 8080
	}

	StringDataBaseConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
