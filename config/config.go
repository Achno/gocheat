package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type ItemWrapper struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
}

type ThemeWrapper struct {
	SubText string `json:"subtext"`
	Accent  string `json:"accent"`
}

type Options struct {
	Items  []ItemWrapper `json:"items"`
	Styles ThemeWrapper  `json:"styles"`
}

// global options object
var GoCheatOptions = Options{}

// Create config if they dont exist and read it into GoCheatOptions
func Init() {

	filepath, err := CreateConfig()

	if err != nil {
		log.Fatalf("Error creating config: %v", err)
	}

	configFile, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer configFile.Close()

	byteValue, _ := io.ReadAll(configFile)
	if err := json.Unmarshal(byteValue, &GoCheatOptions); err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

}

// Ensures ~/.config/gocheat/config.json is created with the template and returns the filepath
func CreateConfig() (string, error) {

	// look for $XDG_CONFIG_HOME/gocheat/config.json or $HOME/.config/gocheat/config.json
	configDir, err := os.UserConfigDir()

	if err != nil {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			// cant find home or config just give up
			return "", fmt.Errorf("while trying to create config: %w ", err)
		}
		configDir = filepath.Join(homeDir, ".config")
	}

	// Create ~/.config/gocheat/ if it does not exist
	configPath := filepath.Join(configDir, "gocheat")
	err = os.MkdirAll(configPath, 0755)

	if err != nil {
		return "", err
	}

	// ~/.config/gocheat/config.json
	configFilePath := filepath.Join(configPath, configFile)

	// check if ~/.config/gocheat/config.json exists
	if _, err = os.Stat(configFilePath); errors.Is(err, os.ErrNotExist) {

		// if not create it with the template.json
		file, err := os.Create(configFilePath)
		if err != nil {
			return "", err
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")

		if err := encoder.Encode(jsonTemplate); err != nil {
			return "", fmt.Errorf("while encoding the template.json : %w", err)
		}

	}

	return configFilePath, nil

}

// updates config.json with the new Values of the GoCheatOptions config object
func UpdateConfig() error {

	updatedConfig, err := json.MarshalIndent(GoCheatOptions, "", "  ")
	if err != nil {
		return fmt.Errorf("error adding Item to config.json: %w", err)
	}

	configFilePath, err := CreateConfig()

	if err != nil {
		return err
	}

	// Save the updated config to config.json
	err = os.WriteFile(configFilePath, updatedConfig, 0644)
	if err != nil {
		return err
	}

	return nil

}
