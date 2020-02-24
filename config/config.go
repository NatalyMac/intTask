package config

import (
	"os"

	"github.com/creasty/defaults"
)

// GetConfig returns config struct
func GetConfig() *Config {
	defaultConfig := &Config{}
	if err := defaults.Set(defaultConfig); err != nil {
		panic(err)
	}
	return &Config{
		ArticlesURL:   getEnv("A_URL", defaultConfig.ArticlesURL),
		CMArticlesURL: getEnv("CM_URL", defaultConfig.CMArticlesURL),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Config describes .env file and set default values
type Config struct {
	ArticlesURL   string `default:"https://storage.googleapis.com/aller-structure-task/articles.json"`
	CMArticlesURL string `default:"https://storage.googleapis.com/aller-structure-task/contentmarketing.json"`
}
