package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port        string
	Hostname    string
	BoringWords []string
}

func BuildConfig() Config {
	homePath := os.Getenv("GO_HOME")

	file, err := os.Open(homePath + "config.json")
	if err != nil {
		fmt.Println(homePath + "config.json")
		fmt.Println("No config file found")
		return DefaultConfig()
	}

	config := Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
		return DefaultConfig()
	}

	return config
}

func DefaultConfig() Config {
	return Config{
		Port:        "1400",
		Hostname:    "127.0.0.1",
		BoringWords: []string{""},
	}
}

func addEnvVariables(c *Config) {
	if host := os.Getenv("GO_HOSTNAME"); host != "" {
		c.Hostname = host
	}

	if port := os.Getenv("GO_PORT"); port != "" {
		c.Port = port
	}
}
