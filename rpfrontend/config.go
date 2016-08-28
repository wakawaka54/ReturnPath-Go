/*
Reads config/config.json and sets up a Config structure to pass these values to caller
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var HomePath string

type Config struct {
	Port       string
	Hostname   string
	ApiAddress string
}

func BuildConfig() Config {
	HomePath = os.Getenv("GO_HOME")
	file, err := os.Open(HomePath + "config.json")
	if err != nil {
		fmt.Println("No config file found, reverting to default config")
		return DefaultConfig()
	}

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Reverting to default config")
		return DefaultConfig()
	}

	addEnvVariables(&config)

	return config
}

func DefaultConfig() Config {
	return Config{
		Port:       "1500",
		Hostname:   "127.0.0.1",
		ApiAddress: "http://localhost:1337",
	}
}

func addEnvVariables(c *Config) {
	if host := os.Getenv("GO_HOSTNAME"); host != "" {
		c.Hostname = host
	}

	if port := os.Getenv("GO_PORT"); port != "" {
		c.Port = port
	}

	if apiaddress := os.Getenv("GO_APIADDRESS"); apiaddress != "" {
		c.ApiAddress = apiaddress
	}
}
