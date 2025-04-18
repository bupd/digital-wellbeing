package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config structure to match the TOML structure
type Config struct {
	PORT string `mapstructure:"PORT"`
}

// Helper function to read the config file using Viper
func ReadConfig() (*viper.Viper, error) {
	dirname := GetHomeDir()
	configPath := fmt.Sprintf("%s/.config/digital-wellbeing/config.toml", dirname)
	// Initialize a new Viper instance
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("toml")

	// Read the configuration file
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w. Please ensure the config file exists.", err)
	}

	return v, nil
}

// Reads config and gives the config
func GetConfig() Config {
	// Load the config file using Viper
	v, err := ReadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Map the values from Viper into the Config struct
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	// Now use the loaded config values
	fmt.Printf("PORT: %s\n", config.PORT)

	return config
}

func GetHomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("unable to get user home directory: %v", err)
	}
	// fmt.Println(dirname)

	return dirname
}
