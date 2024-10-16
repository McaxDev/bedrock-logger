package main

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port       string
	Password   string
	DB         DbConfig
	ChatLength int
}

type DbConfig struct {
	Type     string
	Host     string
	Port     string
	DB       string
	Username string
	Password string
}

var config Config

func GetConfig() {
	config.Port = os.Getenv("PORT")
	config.Password = os.Getenv("PASSWORD")
	config.DB.Type = os.Getenv("DB_TYPE")
	config.DB.Host = os.Getenv("DB_HOST")
	config.DB.Port = os.Getenv("DB_PORT")
	config.DB.DB = os.Getenv("DB_DB")
	config.DB.Username = os.Getenv("DB_USERNAME")
	config.DB.Password = os.Getenv("DB_PASSWORD")

	LoadIntConfig("CHAT_LENGTH", &config.ChatLength, 10)
}

func LoadIntConfig(env string, config *int, defaultValue int) {
	rawValue, exists := os.LookupEnv(env)
	if !exists {
		*config = defaultValue
	} else {
		value, err := strconv.Atoi(rawValue)
		if err != nil {
			log.Fatalf("Invalid %s: %v\n", env, err)
			return
		}
		*config = value
	}
}
