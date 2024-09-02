package config

import (
	"github.com/spf13/viper"
	"log"
)

// Config holds the configuration values for the application.
type Config struct {
	DatabaseURL string
	KafkaBroker string
	JWTSecret   string
	Port        string
}

var Conf Config

// LoadConfig loads configuration from environment variables or configuration files.
func LoadConfig() {
	// Set up Viper to read environment variables or a .env file
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Set default values if environment variables are not set
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("DATABASE_URL", "postgres://user:password@localhost:5432/company_db?sslmode=disable")
	viper.SetDefault("KAFKA_BROKER", "localhost:9092")
	viper.SetDefault("JWT_SECRET", "Gehheim1310")

	// Read the configuration file, if it exists
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No .env file found, using environment variables")
	}

	// Populate the Config struct with the configuration values
	Conf = Config{
		Port:        viper.GetString("PORT"),
		DatabaseURL: viper.GetString("DATABASE_URL"),
		KafkaBroker: viper.GetString("KAFKA_BROKER"),
		JWTSecret:   viper.GetString("JWT_SECRET"),
	}
}
