package config

import "os"

type LogLevel string

const (
	Debug  LogLevel = "debug"
	Info   LogLevel = "info"
	Notice LogLevel = "notice"
	Warn   LogLevel = "warn"
	Error  LogLevel = "error"
)

func GetLogLevel() LogLevel {
	if isDebug() {
		return Debug
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		return Info
	}

	return LogLevel(logLevel)
}

func isDebug() bool {
	return os.Getenv("DEBUG") == "true"
}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}
