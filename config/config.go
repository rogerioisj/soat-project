package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	Port     string `json:"port"`
	Host     string `json:"host"`
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"database"`
}

func (c *Configuration) Validate() error {
	if c.Port == "" {
		return fmt.Errorf("invalid port: %v", c.Port)
	}

	if c.Host == "" {
		return fmt.Errorf("invalid host: %s", c.Host)
	}

	if c.Database.Host == "" {
		return fmt.Errorf("invalid database host: %s", c.Database.Host)
	}

	if c.Database.Port == "" {
		return fmt.Errorf("invalid database port: %v", c.Database.Port)
	}

	if c.Database.User == "" {
		return fmt.Errorf("invalid database user: %s", c.Database.User)
	}

	if c.Database.Password == "" {
		return fmt.Errorf("invalid database password: %s", c.Database.Password)
	}

	if c.Database.Database == "" {
		return fmt.Errorf("invalid database name: %s", c.Database.Database)
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
	config.Database.Host = os.Getenv("DB_HOST")
	config.Database.Port = os.Getenv("DB_PORT")
	config.Database.User = os.Getenv("DB_USER")
	config.Database.Password = os.Getenv("DB_PASSWORD")
	config.Database.Database = os.Getenv("DB_NAME")

	if err := config.Validate(); err != nil {
		return err
	}

	return nil
}
