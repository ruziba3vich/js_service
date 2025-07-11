package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		AppPort, JsContainerName, ContainerWorkingDir, JsSourceFileName, JsExecutableName, LogPath string
		CompilationTimeout, ExecutionTimeout                                                       time.Duration
	}
)

func NewConfig() *Config {
	_ = godotenv.Load()
	return &Config{
		LogPath:             getEnv("LOG_PATH", "app.log"),
		AppPort:             getEnv("APP_PORT", "704"),
		JsContainerName:     getEnv("JS_CONTAINER_NAME", "js-executor"),
		JsExecutableName:    getEnv("JS_EXECUTABLE_NAME", "program"),
		JsSourceFileName:    getEnv("JS_SOURCE_FILE_NAME", "main.js"),
		ContainerWorkingDir: getEnv("CONTAINER_WORKING_DIR", "/app"),
		CompilationTimeout:  getEnvTime("COMPILATION_TIME_OUT", 15),
		ExecutionTimeout:    getEnvTime("EXECUTION_TIME_OUT", 60),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvTime(key string, fallback int) time.Duration {
	if value := os.Getenv(key); value != "" {
		v, err := strconv.Atoi(value)
		if err == nil {
			return time.Duration(v) * time.Second
		}

	}
	return time.Duration(fallback) * time.Second
}
