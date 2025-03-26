package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const CONFIG_FILE_NAME = ".git-profiles-config.json"

type Profile struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	AbsoluteSshPath string `json:"absoluteSshPath"`
}

type Config struct {
	Profiles []Profile `json:"profiles"`
}

func GetConfigPath() string {
	userHomeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatalf("Failed to get the user homedir %v\n", err)
		panic(err)
	}

	return fmt.Sprintf("%s/%s", userHomeDir, CONFIG_FILE_NAME)
}

func FindOrCreateConfigFile() Config {
	configPath := GetConfigPath()

	if _, err := os.Stat(configPath); err != nil {
		CreateConfigFile()
	}

	var config, err = FindConfigFile()

	if err != nil {
		log.Fatalf("Failed to find a config file %v\n", config)
		panic(err)
	}

	return config
}

func FindConfigFile() (Config, error) {
	var file, err = os.ReadFile(GetConfigPath())

	if err != nil {
		return Config{}, err
	}

	var config Config

	err = json.Unmarshal(file, config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func CreateConfigFile() {
	var emptyConfig = Config{
		Profiles: make([]Profile, 0),
	}

	var jsonConfig, err = json.Marshal(emptyConfig)

	if err != nil {
		log.Fatalf("Failed to create a config file %v\n", err)
		panic(err)
	}

	err = os.WriteFile(GetConfigPath(), jsonConfig, 0644)

	if err != nil {
		log.Fatalf("Failed to create config file %v\n", err)
		panic(err)
	}
}
