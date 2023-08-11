package aliases

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var configDir = "config"
var configFileName = ".mockrc" // Specify only the file name here

func ExecuteLoadAliases() (map[string]string, error) {
	// Check if the config directory exists, and create it if not
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.Mkdir(configDir, os.ModePerm)
		if err != nil {
			return nil, err
		}

		// Generate default config
		if err := generateDefaultConfig(); err != nil {
			return nil, err
		}
	}

	// Use filepath.Join to create the full file path
	configFile, err := os.Open(filepath.Join(configDir, configFileName))
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config struct {
		Aliases map[string]string `json:"aliases"`
	}

	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return nil, err
	}

	return config.Aliases, nil
}

func generateDefaultConfig() error {
	defaultConfig := struct {
		Aliases map[string]string `json:"aliases"`
	}{
		Aliases: map[string]string{
			"dir": "ls",
			"diskusage": "du",
			"clear": "cls",
		},
	}

	configFile, err := os.Create(filepath.Join(configDir, configFileName))
	if err != nil {
		return err
	}
	defer configFile.Close()

	encoder := json.NewEncoder(configFile)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(&defaultConfig); err != nil {
		return err
	}

	fmt.Println("Initialized default config:", configFileName)
	return nil
}

func SaveAliases(aliases map[string]string) error {
	// Check if the config directory exists, and create it if not
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.Mkdir(configDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	configFile, err := os.Create(filepath.Join(configDir, configFileName))
	if err != nil {
		return err
	}
	defer configFile.Close()

	config := struct {
		Aliases map[string]string `json:"aliases"`
	}{
		Aliases: aliases,
	}

	encoder := json.NewEncoder(configFile)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(&config); err != nil {
		return err
	}

	return nil
}
