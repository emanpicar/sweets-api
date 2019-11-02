package settings

import (
	"os"
)

func getEnv(envName, envDefault string) string {
	if envValue := os.Getenv(envName); envValue != "" {
		return envValue
	}

	return envDefault
}

func GetLogLevel() string {
	return getEnv("LOG_LEVEL", "info")
}
