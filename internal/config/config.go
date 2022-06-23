package config

import (
	"log"

	"github.com/joho/godotenv"
)

var (
	CORSSupport      bool
	FrontendHostname string
	PublishAddress   string
)

func init() {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	errEnv := godotenv.Load(".env")

	if errEnv != nil {
		log.Fatalf("Error loading .env file")
	}
	// CORSSupport
	CORSSupport = get(EnvCORSSupport) == "yes"
	// PublishAddress
	PublishAddress = get(EnvPublishAddress)
}
