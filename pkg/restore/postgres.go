package restore

import (
	"fmt"
	"os/exec"
)


func RestorePostgres(host string, port int, user, password, dbName, input string) error {
	cmdStr := fmt.Sprintf("PGPASSWORD=%s pg_restore -h %s -p %d -U %s -d %s -v %s",
		password, host, port, user, dbName, input)
	cmd := exec.Command("sh", "-c", cmdStr)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("pg_restore error: %v, output: %s", err, string(out))
	}
	return nil
}
