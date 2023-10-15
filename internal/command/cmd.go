package command

import (
	"log/slog"
	"os/exec"
)

func RunCmd(cmd string, args ...string) ([]byte, error) {
	c := exec.Command(cmd, args...)
	slog.Info("exec command", "command", c.String())
	return c.CombinedOutput()
}
