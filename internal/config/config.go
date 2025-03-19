package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds all configuration for our application
type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	JWTSecret  string
	ServerPort string
}

// LoadConfig loads the configuration from .env file and environment variables
func LoadConfig() *Config {
	// Set up Viper to read from .env file
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("env")  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	viper.AutomaticEnv()

	// Attempt to read the config file
	if err := viper.ReadInConfig(); err != nil {
		// It's okay if .env file doesn't exist
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Printf("Error reading config file: %v\n", err)
		}
	}

	// Set default values
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASSWORD", "abiobi")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("DB_NAME", "inventory")
	viper.SetDefault("JWT_SECRET", "your-secret-key")
	viper.SetDefault("SERVER_PORT", "8080")

	// Create the config
	return &Config{
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBName:     viper.GetString("DB_NAME"),
		JWTSecret:  viper.GetString("JWT_SECRET"),
		ServerPort: viper.GetString("SERVER_PORT"),
	}
}