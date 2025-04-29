package utils

import (
	"fmt"
	"log"
	"os/exec"
)

func RunIpmiCommand(args ...string) error {
	log.Printf("Executing ipmitool command: ipmitool %v", args)
	cmd := exec.Command("ipmitool", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Command failed with error: %v, output: %s", err, string(output))
		return fmt.Errorf("failed to run ipmitool command: %w, output: %s", err, string(output))
	}
	log.Printf("Command succeeded with output: %s", string(output))
	return nil
}
