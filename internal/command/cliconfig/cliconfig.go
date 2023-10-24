package cliconfig

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type BuildConfig struct {
	IgnoreFiles []string
}

type Config struct {
	Build      BuildConfig
	Title      string
	Descrition string
	BaseUrl    string
}

// LoadConfig loads a TOML text into a Config struct.
func LoadConfig() (*Config, error) {
	path, err := cliConfigFile()
	if err != nil {
		return nil, err
	}
	return loadConfigFile(path)
}

func loadConfigFile(path string) (*Config, error) {
	log.Printf("[INFO] Loading config file: %s", path)
	d, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = toml.Unmarshal(d, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// AsMap returns a map[string]interface{} representation of the Config struct.
func (c *Config) AsMap() map[string]interface{} {
	return map[string]interface{}{
		"build": map[string]interface{}{
			"ignore_files": c.Build.IgnoreFiles,
		},
		"title":       c.Title,
		"description": c.Descrition,
		"base_url":    c.BaseUrl,
	}
}

func configFile() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.toml"), nil
}

// cliConfigFileOverride returns the value of VSS_CONFIG_FILE if set.
func cliConfigFileOverride() string {
	return os.Getenv("VSS_CONFIG_FILE")
}

func cliConfigFile() (string, error) {
	configFilePath := cliConfigFileOverride()
	if configFilePath == "" {
		var err error
		configFilePath, err = configFile()
		if err != nil {
			return "", err
		}
	}

	f, err := os.Open(configFilePath)
	if err == nil {
		f.Close()
		return configFilePath, nil
	} else {
		return "", err
	}
}
