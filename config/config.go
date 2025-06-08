package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	Port        string `json:"port"`
	Host        string `json:"host"`
	DatabaseUrl string `json:"database_url"`
}

func (c *Configuration) Validate() error {
	if c.Port == "" {
		return fmt.Errorf("invalid port: %v", c.Port)
	}

	if c.Host == "" {
		return fmt.Errorf("invalid host: %s", c.Host)
	}

	if c.DatabaseUrl == "" {
		return fmt.Errorf("invalid database host: %s", c.DatabaseUrl)
	}

	return nil
}

func Load() (Configuration, error) {
	environment := os.Getenv("PROJECT_ENV")

	var config Configuration

	switch environment {
	case "production":
		log.Println("Loading production configuration")
		err := loadEnvironment(&config)
		return config, err
	default:
		log.Println("Loading development configuration")
		err := loadDevEnvironmentVars(&config)
		return config, err
	}
}

func loadDevEnvironmentVars(config *Configuration) error {
	file, err := os.Open("env.json")
	if err != nil {
		return fmt.Errorf("could not open env file: %v", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)

	if err != nil {
		return fmt.Errorf("could not parse env file: %v", err)
	}

	err = config.Validate()

	if err != nil {
		return err
	}

	return nil
}

func loadEnvironment(config *Configuration) error {

	config.Host = os.Getenv("HOST")
	config.Port = os.Getenv("PORT")
	config.DatabaseUrl = os.Getenv("DATABASE_URL")

	if err := config.Validate(); err != nil {
		return err
	}

	return nil
}
