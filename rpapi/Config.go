/*
Adds Config.json functionality, also reads environment variables.

Environment variables take presidence over config.json values.

BuildConfig() returns a Config object with configuration values

ENV_VARS

GO_HOME = Path to config.json, usually the project directory, not needed if
					binary is started from within the project folder
GO_PORT = Listening port
GO_HOSTNAME = Listening IP Address, use 0.0.0.0 to avoid localhost, 127.0.0.1 problems

Config.json
Port = Listening Port
Hostname = Listening IP Address
BoringWords = Words to ignore when adding sentences
*/

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
