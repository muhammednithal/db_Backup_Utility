package config

import (
	"encoding/json"
	"errors"
	"os"
)

const configFilePath = "config.json"

type DBConfig struct {
	DBType string `json:"dbType"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	DBName string `json:"dbName"`
	Output string `json:"output"` // Used as input in restore
}

func LoadConfigFile() map[string]DBConfig {
	configMap := make(map[string]DBConfig)
	if _, err := os.Stat(configFilePath); err == nil {
		data, err := os.ReadFile(configFilePath)
		if err == nil {
			json.Unmarshal(data, &configMap)
		}
	}
	return configMap
}

func SaveConfigFile(configMap map[string]DBConfig) error {
	data, err := json.MarshalIndent(configMap, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configFilePath, data, 0644)
}

func DeleteConfig(variant string) error {
	configMap := LoadConfigFile()
	if _, ok := configMap[variant]; !ok {
		return errors.New("config not found")
	}
	delete(configMap, variant)
	return SaveConfigFile(configMap)
}

func SaveVariant(variant string, config DBConfig) error {
	configMap := LoadConfigFile()
	configMap[variant] = config
	return SaveConfigFile(configMap)
}

func GetVariant(variant string) (DBConfig, error) {
	configMap := LoadConfigFile()
	config, ok := configMap[variant]
	if !ok {
		return DBConfig{}, errors.New("config variant not found")
	}
	return config, nil
}
