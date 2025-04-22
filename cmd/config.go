/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

const configFilePath = "config.json"

type DBConfig struct {
	DBType string `json:"dbType"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	DBName string `json:"dbName"`
	Output string `json:"output"`
}
// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage dbtool configuration settings",
	Long: `Allows you to view, edit, or reset dbtool's configuration file.
This includes database connection settings, backup locations, and other preferences.`,
	Run: func(cmd *cobra.Command, args []string) {
		var qs = []*survey.Question{
			{
				Name: "dbType",
				Prompt: &survey.Select{
					Message: "Choose database type:",
					Options: []string{"mysql", "postgres"},
				},
				Validate: survey.Required,
			},
			{
				Name:     "host",
				Prompt:   &survey.Input{Message: "Database host:", Default: "localhost"},
				Validate: survey.Required,
			},
			{
				Name:     "port",
				Prompt:   &survey.Input{Message: "Database port:", Default: "3306"},
				Validate: survey.Required,
			},
			{
				Name:     "user",
				Prompt:   &survey.Input{Message: "Database username:"},
				Validate: survey.Required,
			},
			{
				Name:     "dbName",
				Prompt:   &survey.Input{Message: "Database name:"},
				Validate: survey.Required,
			},
			{
				Name:     "output",
				Prompt:   &survey.Input{Message: "Output file path:", Default: "backup.sql"},
				Validate: survey.Required,
			},
		}

		answers := struct {
			DBType string `survey:"dbType"`
			Host   string `survey:"host"`
			Port   string `survey:"port"`
			User   string `survey:"user"`
			DBName string `survey:"dbName"`
			Output string `survey:"output"`
		}{}

		if err := survey.Ask(qs, &answers); err != nil {
			fmt.Println("Prompt failed:", err)
			return
		}

		var variantName string
		survey.AskOne(&survey.Input{
			Message: "Config variant name:",
		}, &variantName, survey.WithValidator(survey.Required))

		portInt, _ := strconv.Atoi(answers.Port)

		newConfig := DBConfig{
			DBType: answers.DBType,
			Host:   answers.Host,
			Port:   portInt,
			User:   answers.User,
			DBName: answers.DBName,
			Output: answers.Output,
		}

		configMap := loadConfigFile()
		configMap[variantName] = newConfig

		data, _ := json.MarshalIndent(configMap, "", "  ")
		os.WriteFile(configFilePath, data, 0644)

		fmt.Println("Configuration saved under variant:", variantName)
	},
}

func loadConfigFile() map[string]DBConfig {
	configMap := make(map[string]DBConfig)

	if _, err := os.Stat(configFilePath); err == nil {
		data, err := os.ReadFile(configFilePath)
		if err == nil {
			json.Unmarshal(data, &configMap)
		}
	}
	return configMap
}
func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
