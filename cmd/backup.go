/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/muhammednithal/db_Backup_Utility/pkg/backup"
	"github.com/spf13/cobra"
)


var (
	dbType, host, user, password, dbName, output string
	port                                         int
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Create a backup of the database",
	Long: `Creates a backup of the configured database and stores it in the specified location.
Supports options like output directory, compression, and custom filenames.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch dbType {
		case "mysql":
			err := backup.BackupMySQL(host, port, user, password, dbName, output)
			handleBackupResult(err)
		case "postgres":
			err := backup.BackupPostgres(host, port, user, password, dbName, output)
			handleBackupResult(err)
		default:
			fmt.Printf("Unsupported database type: %s\n", dbType)
		}
	},
}

func handleBackupResult(err error) {
	if err != nil {
		fmt.Println("Backup failed:", err)
	} else {
		fmt.Println("Backup completed successfully.")
	}
}



func init() {
	rootCmd.AddCommand(backupCmd)

	backupCmd.Flags().StringVar(&dbType, "type", "", "Database type (mysql)")
	backupCmd.Flags().StringVar(&host, "host", "localhost", "Database host")
	backupCmd.Flags().IntVar(&port, "port", 3306, "Database port")
	backupCmd.Flags().StringVar(&user, "user", "", "Database username")
	backupCmd.Flags().StringVar(&password, "password", "", "Database password")
	backupCmd.Flags().StringVar(&dbName, "database", "", "Database name")
	backupCmd.Flags().StringVar(&output, "output", "backup.sql", "Output file path for backup")

	backupCmd.MarkFlagRequired("type")
	backupCmd.MarkFlagRequired("user")
	backupCmd.MarkFlagRequired("password")
	backupCmd.MarkFlagRequired("database")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
