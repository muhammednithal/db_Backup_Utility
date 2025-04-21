package backup

import (
	"fmt"
	"os/exec"
)

func BackupMySQL(host string, port int, user, password, dbName, output string) error {
	cmdStr := fmt.Sprintf("mysqldump -h %s -P %d -u %s -p%s %s > %s", host, port, user, password, dbName, output)
	outputBytes, err := exec.Command("sh", "-c", cmdStr).CombinedOutput()
	if err != nil {
		return fmt.Errorf("mysqldump error: %v, output: %s", err, string(outputBytes))
	}
	return nil
}
