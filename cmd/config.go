package cmd

import (
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	cfg "github.com/muhammednithal/db_Backup_Utility/pkg/config"
	"github.com/spf13/cobra"
)

var deleteVariant string

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage dbtool configuration settings",
	Long:  `View, edit, delete, or create new configuration variants.`,
	Run: func(cmd *cobra.Command, args []string) {
		if deleteVariant != "" {
			if err := cfg.DeleteConfig(deleteVariant); err != nil {
				fmt.Println("Error deleting config:", err)
			} else {
				fmt.Println("Deleted config variant:", deleteVariant)
			}
			return
		}

		// Prompt for new config
		var answers struct {
			DBType string `survey:"dbType"`
			Host   string `survey:"host"`
			Port   string `survey:"port"`
			User   string `survey:"user"`
			DBName string `survey:"dbName"`
			Output string `survey:"output"`
		}

		qs := []*survey.Question{
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

		if err := survey.Ask(qs, &answers); err != nil {
			fmt.Println("Prompt failed:", err)
			return
		}

		var variantName string
		survey.AskOne(&survey.Input{Message: "Config variant name:"}, &variantName, survey.WithValidator(survey.Required))

		portInt, _ := strconv.Atoi(answers.Port)

		err := cfg.SaveVariant(variantName, cfg.DBConfig{
			DBType: answers.DBType,
			Host:   answers.Host,
			Port:   portInt,
			User:   answers.User,
			DBName: answers.DBName,
			Output: answers.Output,
		})
		if err != nil {
			fmt.Println("Failed to save config:", err)
		} else {
			fmt.Println("Config saved as:", variantName)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVar(&deleteVariant, "delete", "", "Delete a saved config variant by name")
}
