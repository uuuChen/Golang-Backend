package config

import (
	"os"
)

type Config struct {
	MySQLRootPassword string
	MySQLDatabase     string
	MySQLUser         string
	MySQLPassword     string
	RedisHost         string
	RedisPort         string
	JWTSecret         []byte
}

var AppConfig *Config

func LoadConfigFromEnvFile() {
	AppConfig = &Config{
		MySQLRootPassword: getEnv("MYSQL_ROOT_PASSWORD", "defaultRootPassword"),
		MySQLDatabase:     getEnv("MYSQL_DATABASE", "defaultDatabase"),
		MySQLUser:         getEnv("MYSQL_USER", "defaultUser"),
		MySQLPassword:     getEnv("MYSQL_PASSWORD", "defaultPassword"),
		RedisHost:         getEnv("REDIS_HOST", "redis"),
		RedisPort:         getEnv("REDIS_PORT", "6379"),
		JWTSecret:         []byte(getEnv("JWT_SECRET", "defaultSecret")),
	}
}

// Helper function to read an environment variable or return a default value if not set
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
