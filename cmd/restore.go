/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/muhammednithal/db_Backup_Utility/pkg/restore"
	"github.com/spf13/cobra"
)

var input string
// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a database from a backup file",
	Long: `Restores a database using a previously created backup file.
You can specify the path to the backup file and the target database configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		if dbType == "mysql" {
			err := restore.RestoreMYSQL(host, port, user, password, dbName, input)

			if err != nil {
				fmt.Println("Restore failed:", err)
			} else {
				fmt.Println("Restore completed successfully.")
			}
		} else {
			fmt.Printf("Unsupported database type: %s\n", dbType)
		}
	},
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

	restoreCmd.MarkFlagRequired("type")
	restoreCmd.MarkFlagRequired("user")
	restoreCmd.MarkFlagRequired("password")
	restoreCmd.MarkFlagRequired("database")
	restoreCmd.MarkFlagRequired("input")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restoreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restoreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
