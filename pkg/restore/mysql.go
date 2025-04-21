package restore

import (
	"fmt"
	"os/exec"
)

func RestoreMYSQL(host string, port int, user, password, dbName, input string) error {
	restoreCmd := fmt.Sprintf("mysql -h %s -P %d -u %s -p%s %s < %s", host, port, user, password, dbName, input)
	out, err := exec.Command("sh", "-c", restoreCmd).CombinedOutput()
	if err != nil {
		return fmt.Errorf("restore error: %v, output: %s", err, string(out))
	}
	return nil
}
