package config

import (
	"os"
	"strconv"
)

func GetEnvString(key, fallback string) string {
	if env, ok := os.LookupEnv(key); ok {
		return env
	}

	return fallback
}

func GetEnvInt(key string, fallback int) int {
	env, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	if parsed, err := strconv.Atoi(env); err == nil {
		return parsed
	}

	return fallback
}
