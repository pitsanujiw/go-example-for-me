package config

import (
	"os"
)

func get(key EnvKey) string {
	value := os.Getenv(string(key))

	return value
}
