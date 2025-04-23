package cmd

import (
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	configPkg "github.com/muhammednithal/db_Backup_Utility/pkg/config"
	"github.com/muhammednithal/db_Backup_Utility/pkg/logger"
	"github.com/muhammednithal/db_Backup_Utility/pkg/restore"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a database from a backup file",
	Long: `Restores a database using a previously created backup file.
You can specify the path to the backup file and the target database configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Mode 3: Using a saved config
		if savedConfig != "" {
			config, err := configPkg.GetVariant(savedConfig)
			if err != nil {
				fmt.Println("Error loading config:", err)
				return
			}

			survey.AskOne(&survey.Password{Message: "Enter database password:"}, &password, survey.WithValidator(survey.Required))
			dispatchRestore(config.DBType, config.Host, config.Port, config.User, password, config.DBName, config.Output)
			return
		}

		// Mode 1: Use flags directly if provided
		flagUsed := cmd.Flags().Changed("type") || cmd.Flags().Changed("host") || cmd.Flags().Changed("user") ||
			cmd.Flags().Changed("password") || cmd.Flags().Changed("database")

		if flagUsed {
			dispatchRestore(dbType, host, port, user, password, dbName, input)
			return
		}

		// Mode 2: Prompt user interactively
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
				Prompt:   &survey.Input{Message: "Database user:"},
				Validate: survey.Required,
			},
			{
				Name:     "password",
				Prompt:   &survey.Password{Message: "Database password:"},
				Validate: survey.Required,
			},
			{
				Name:     "dbName",
				Prompt:   &survey.Input{Message: "Database name:"},
				Validate: survey.Required,
			},
			{
				Name:     "input",
				Prompt:   &survey.Input{Message: "Input file path:", Default: "backup.sql"},
				Validate: survey.Required,
			},
		}

		answers := struct {
			DBType   string `survey:"dbType"`
			Host     string `survey:"host"`
			Port     string `survey:"port"`
			User     string `survey:"user"`
			Password string `survey:"password"`
			DBName   string `survey:"dbName"`
			Input    string `survey:"input"`
		}{}

		if err := survey.Ask(qs, &answers); err != nil {
			fmt.Println("Prompt failed:", err)
			return
		}

		portInt, err := strconv.Atoi(answers.Port)
		if err != nil {
			fmt.Println("Invalid port number.")
			return
		}

		dispatchRestore(answers.DBType, answers.Host, portInt, answers.User, answers.Password, answers.DBName, answers.Input)
	},
}

func dispatchRestore(dbType, host string, port int, user, password, dbName, input string) {
	switch dbType {
	case "mysql":
		err := restore.RestoreMYSQL(host, port, user, password, dbName, input)
		handleRestoreResult(err, dbType, host, port, user, dbName, input)
	case "postgres":
		err := restore.RestorePostgres(host, port, user, password, dbName, input)
		handleRestoreResult(err, dbType, host, port, user, dbName, input)
	default:
		fmt.Printf("Unsupported database type: %s\n", dbType)
	}
}

func handleRestoreResult(err error,dbType, host string, port int, user, dbName, input string) {
	status := "success"
	errorMsg := ""
	if err != nil {
		status = "failure"
		errorMsg = err.Error()
		fmt.Println("Restore failed:", err)
	} else {
		fmt.Println("Restore completed successfully.")
	}

	logger.LogOperation(logger.LogEntry{
		Action:      "restore",
		DBType:      dbType,
		Host:        host,
		Port:        port,
		User:        user,
		Database:    dbName,
		FilePath:    input,
		Status:      status,
		Error:       errorMsg,
		SavedConfig: savedConfig,
	})
}


func init() {
	rootCmd.AddCommand(restoreCmd)

	restoreCmd.Flags().StringVar(&dbType, "type", "", "Database type (mysql|postgres)")
	restoreCmd.Flags().StringVar(&host, "host", "localhost", "Database host")
	restoreCmd.Flags().IntVar(&port, "port", 3306, "Database port")
	restoreCmd.Flags().StringVar(&user, "user", "", "Database username")
	restoreCmd.Flags().StringVar(&password, "password", "", "Database password")
	restoreCmd.Flags().StringVar(&dbName, "database", "", "Database name")
	restoreCmd.Flags().StringVar(&input, "input", "", "Input backup file path")
	restoreCmd.Flags().StringVar(&savedConfig, "savedconfig", "", "Use a saved config (provide config name)")
}
