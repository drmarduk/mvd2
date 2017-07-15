package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config for the api endpoints
type Config struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	TemplateDir string `json:"templatedir"`

	DBHost     string `json:"db_host,omitempty"`
	DBDatabase string `json:"db_database,omitempty"`
	DBUser     string `json:"db_user,omitempty"`
	DBPassword string `json:"db_password,omitempty"`
	DBDriver   string `json:"db_driver,omitempty"`
}

// HTTPAddress formats host and port to a combined string
func (c *Config) HTTPAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// NewConfig returns a json object from the config.json file or error
func NewConfig(configfile string) (*Config, error) {
	src, err := ioutil.ReadFile(configfile)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(src, config)

	return config, err
}
