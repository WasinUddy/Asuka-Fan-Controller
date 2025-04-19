package utils

import (
	"fmt"
	"os/exec"
)

func RunIpmiCommand(args ...string) error {
	cmd := exec.Command("ipmitool", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run ipmitool command: %w, output: %s", err, string(output))
	}
	return nil
}
