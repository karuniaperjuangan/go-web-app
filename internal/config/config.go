package config

import (
	"log"
	"os"
)

type Config struct {
	JWTSecret string
}

func LoadConfig() *Config {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		log.Fatalln("SECRET KEY NOT SET! EXITING")
	}
	return &Config{
		JWTSecret: secret,
	}
}
