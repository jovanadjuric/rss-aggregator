package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var contents Config

	err = json.Unmarshal(file, &contents)
	if err != nil {
		return Config{}, err
	}

	return contents, nil
}

func (c Config) SetUser(username string) error {
	c.Current_User_Name = &username

	err := write(c)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homedir + "/" + configFileName, nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	contents, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, contents, 0644)
	if err != nil {
		return err
	}

	return nil
}
