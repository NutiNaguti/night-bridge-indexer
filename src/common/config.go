package common

import "os"

type DbConfig struct {
	ConnectionString string
}

type Config struct {
	Db DbConfig
}

func New() *Config {
	return &Config{
		Db: DbConfig{
			ConnectionString: getEnv("DB_CONNECTION_STRING", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
