package backup

import (
	"fmt"
	"os/exec"
)

func BackupPostgres(host string, port int, user, password, dbName, output string) error {
	cmdStr := fmt.Sprintf("PGPASSWORD=%s pg_dump -h %s -p %d -U %s -F c -b -v -f %s %s",
		password, host, port, user, output, dbName)
	cmd := exec.Command("sh", "-c", cmdStr)

	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("pg_dump error: %v, output: %s", err, string(outputBytes))
	}
	return nil
}
