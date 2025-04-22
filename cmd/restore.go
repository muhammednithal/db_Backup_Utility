/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/muhammednithal/db_Backup_Utility/pkg/restore"
	"github.com/spf13/cobra"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a database from a backup file",
	Long: `Restores a database using a previously created backup file.
You can specify the path to the backup file and the target database configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		if savedConfig != "" {


            configs := loadConfigFile()
            config, ok := configs[savedConfig]


            if !ok {
                fmt.Println(" No saved config found with name:", savedConfig)
                return
            }
       
            survey.AskOne(&survey.Password{Message: "Enter database password:"}, &password, survey.WithValidator(survey.Required))      
            dispatchRestore(config.DBType, config.Host, config.Port, config.User, password, config.DBName, config.Output)
            return
        }




        // Mode 1: If required flags are set (non-zero), use them directly
        flagUsed := cmd.Flags().Changed("type") || cmd.Flags().Changed("host") || cmd.Flags().Changed("user") ||
            cmd.Flags().Changed("password") || cmd.Flags().Changed("database")




        if flagUsed {
            dispatchRestore(dbType, host, port, user, password, dbName, input)
            return
        }




        // Mode 2: Prompt interactively
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
                Prompt:   &survey.Input{Message: "input file path:", Default: "backup.sql"},
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
            Output   string `survey:"input"`
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




        dispatchRestore(answers.DBType, answers.Host, portInt, answers.User, answers.Password, answers.DBName, answers.Output)
	},
}

func dispatchRestore(dbType, host string, port int, user, password, dbName, input string) {
    switch dbType {
    case "mysql":
        err := restore.RestoreMYSQL(host, port, user, password, dbName, input)
        handleRestoreResult(err)
    case "postgres":
        err := restore.RestorePostgres(host, port, user, password, dbName, input)
        handleRestoreResult(err)
    default:
        fmt.Printf("Unsupported database type: %s\n", dbType)
    }
}

func handleRestoreResult(err error) {
	if err != nil {
		fmt.Println("Restore failed:", err)
	} else {
		fmt.Println("Restore completed successfully.")
	}
}

func init() {
	rootCmd.AddCommand(restoreCmd)

	restoreCmd.Flags().StringVar(&dbType, "type", "", "Database type (mysql)")
    restoreCmd.Flags().StringVar(&host, "host", "localhost", "Database host")
    restoreCmd.Flags().IntVar(&port, "port", 3306, "Database port")
    restoreCmd.Flags().StringVar(&user, "user", "", "Database username")
    restoreCmd.Flags().StringVar(&password, "password", "", "Database password")
    restoreCmd.Flags().StringVar(&dbName, "database", "", "Database name")
    restoreCmd.Flags().StringVar(&input, "input", "", "Input backup file path")
    restoreCmd.Flags().StringVar(&savedConfig, "savedconfig", "", "Use a saved config (provide config name)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restoreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restoreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
