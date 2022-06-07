package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const Default_Config_File_Type = "json"
const Default_Config_Folder = "config"
const Default_Config_File_Name = "dev"

var Default_Path = "config/dev.json"

func InitConfiguration(path string, fileType string) (*Config, error) {

	if fileType == "" {
		fileType = Default_Config_File_Type
	}

	if path != "" {
		Default_Path = fmt.Sprintf("%s.%s", path, fileType)
	}

	config, err := getConfiguration(Default_Path)

	if err != nil {
		return nil, err
	}

	return config, nil
}

func getConfiguration(path string) (*Config, error) {

	var configuration Config

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&configuration)

	if err != nil {
		return nil, err
	}

	return &configuration, nil

}
