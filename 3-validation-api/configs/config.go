package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Email    EmailConfig
	Password PasswordConfig
	Address  AddressConfig
}

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type PasswordConfig struct {
	MinLength int
	MaxLength int
}

type AddressConfig struct {
	MaxLength int
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}

	emailPort, err := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	if err != nil {
		log.Println("Error parsing EMAIL_PORT, using default value")
		emailPort = 587
	}

	passwordMinLength, err := strconv.Atoi(os.Getenv("PASSWORD_MIN_LENGTH"))
	if err != nil {
		log.Println("Error parsing PASSWORD_MIN_LENGTH, using default value")
		passwordMinLength = 8
	}

	passwordMaxLength, err := strconv.Atoi(os.Getenv("PASSWORD_MAX_LENGTH"))
	if err != nil {
		log.Println("Error parsing PASSWORD_MAX_LENGTH, using default value")
		passwordMaxLength = 64
	}

	addressMaxLength, err := strconv.Atoi(os.Getenv("ADDRESS_MAX_LENGTH"))
	if err != nil {
		log.Println("Error parsing ADDRESS_MAX_LENGTH, using default value")
		addressMaxLength = 255
	}

	return &Config{
		Email: EmailConfig{
			Host:     os.Getenv("EMAIL_HOST"),
			Port:     emailPort,
			Username: os.Getenv("EMAIL_USERNAME"),
			Password: os.Getenv("EMAIL_PASSWORD"),
		},
		Password: PasswordConfig{
			MinLength: passwordMinLength,
			MaxLength: passwordMaxLength,
		},
		Address: AddressConfig{
			MaxLength: addressMaxLength,
		},
	}
}
