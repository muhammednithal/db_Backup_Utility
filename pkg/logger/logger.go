package logger

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// LogEntry represents a structured log for backup/restore operations
type LogEntry struct {
	Timestamp   string `json:"timestamp"`
	Action      string `json:"action"`
	DBType      string `json:"db_type"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	User        string `json:"user"`
	Database    string `json:"database"`
	FilePath    string `json:"file_path"`
	Status      string `json:"status"`
	Error       string `json:"error,omitempty"`
	SavedConfig string `json:"saved_config,omitempty"`
}

var (
	logFilePath = "logs/operations.log"
)

// LogOperation logs a backup or restore operation
func LogOperation(entry LogEntry) {
	entry.Timestamp = time.Now().Format(time.RFC3339)

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Failed to open log file:", err)
		return
	}
	defer file.Close()

	logger := log.New(file, "", 0)
	data, err := json.Marshal(entry)
	if err != nil {
		log.Println("Failed to marshal log entry:", err)
		return
	}

	logger.Println(string(data))
}
