package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT                        string
	JWT_SECRET                  string
	JWT_ACCESS_TOKEN_LIFE_TIME  int
	JWT_REFRESH_TOKEN_LIFE_TIME int
}

var Configuration *Config

func loadConfig() {

	if os.Getenv("DOCKER_ENV") != "true" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Failed to load env")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Failed to load PORT ENV")
	}

	jwt_secret := os.Getenv("JWT_SECRET")
	if jwt_secret == "" {
		log.Fatal("Failed to load JWT_SECRET ENV")
	}

	jwt_access_token_life_time := os.Getenv("JWT_ACCESS_TOKEN_LIFE_TIME")
	if jwt_access_token_life_time == "" {
		log.Fatal("Failed to load JWT_ACCESS_TOKEN_LIFE_TIME ENV")
	}

	jwt_refresh_token_life_time := os.Getenv("JWT_REFRESH_TOKEN_LIFE_TIME")
	if jwt_refresh_token_life_time == "" {
		log.Fatal("Failed to load JWT_REFRESH_TOKEN_LIFE_TIME ENV")
	}

	jwt_access_token_life_time_int, err := strconv.Atoi(jwt_access_token_life_time)
	if err != nil {
		log.Fatal("FAILED TO CONVERT JWT_ACCESS_TOKEN_LIFE_TIME INTO INTEGER")
	}

	jwt_refresh_token_life_time_int, err := strconv.Atoi(jwt_refresh_token_life_time)
	if err != nil {
		log.Fatal("FAILED TO CONVERT JWT_REFRESH_TOKEN_LIFE_TIME INTO INTEGER")
	}

	Configuration = &Config{
		PORT:                        port,
		JWT_SECRET:                  jwt_secret,
		JWT_ACCESS_TOKEN_LIFE_TIME:  jwt_access_token_life_time_int,
		JWT_REFRESH_TOKEN_LIFE_TIME: jwt_refresh_token_life_time_int,
	}

}

func GetConfig() *Config {
	if Configuration == nil {
		loadConfig()
	}

	return Configuration
}
