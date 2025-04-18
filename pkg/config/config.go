package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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

// Helper function to create the config file using Viper
func CreateConfig() (*viper.Viper, error) {
	homedir := GetHomeDir()
	configPath := filepath.Join(homedir, ".config", "digital-wellbeing", "config.toml")

	// Check if the directory exists, create it if it doesn't
	if _, err := os.Stat(filepath.Dir(configPath)); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
			return nil, fmt.Errorf("error creating config directory: %w", err)
		}
	}

	// Check if the file exists, create it with default config if it doesn't
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := createDefaultConfigFile(configPath); err != nil {
			return nil, fmt.Errorf("error creating config file: %w", err)
		}
	}

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

// Helper function to create a default config file
func createDefaultConfigFile(configPath string) error {
	// You can customize this function to create a default TOML config file
	// For example:
	configContent := `# Default configuration for Digital Wellbeing
# Add your configuration here
PORT =  "8888"
`

	// Write the content to the file
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		return fmt.Errorf("error writing default config file: %w", err)
	}
	return nil
}

// Reads config and gives the config
func GetConfig() Config {
	// Load the config file using Viper
	v, err := ReadConfig()
	if err != nil {
		log.Printf("error loading config: %v", err)
		v, err = CreateConfig()
		if err != nil {
			log.Fatalf("error creating config: %v", err)
		}
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
