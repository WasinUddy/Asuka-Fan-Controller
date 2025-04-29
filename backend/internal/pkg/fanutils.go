package pkg

import (
	"fmt"
	"log"
	"os/exec"
)

// These functions handle fan mode and speed operations
// Add these if they don't already exist

func GetFanMode() (string, error) {
	// Use IPMI to get the current fan mode
	// For now, just read from a config file or use a simpler approach
	// This is a placeholder - implement based on your actual setup

	cmd := exec.Command("ipmitool", "raw", "0x30", "0x45", "0x00")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get fan mode: %w", err)
	}

	// Parse IPMI output - adjust based on your server's IPMI implementation
	// This is a simplified example
	if len(output) > 0 && output[0] == 0x01 {
		return "auto", nil
	}
	return "manual", nil
}

func SetFanMode(mode string) error {
	var cmd *exec.Cmd

	if mode == "auto" {
		cmd = exec.Command("ipmitool", "raw", "0x30", "0x45", "0x01", "0x01")
	} else if mode == "manual" {
		cmd = exec.Command("ipmitool", "raw", "0x30", "0x45", "0x01", "0x00")
	} else {
		return fmt.Errorf("invalid fan mode: %s", mode)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to set fan mode %s: %w, output: %s", mode, err, string(output))
	}

	log.Printf("Fan mode set to %s", mode)
	return nil
}

func GetFanSpeed() (int, error) {
	// This is a placeholder - implement based on your actual setup
	cmd := exec.Command("ipmitool", "raw", "0x30", "0x45", "0x00")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("failed to get fan speed: %w", err)
	}

	// Parse IPMI output - adjust based on your server's IPMI implementation
	// This is a simplified example
	if len(output) > 1 {
		return int(output[1]), nil
	}

	return 0, fmt.Errorf("unable to parse fan speed from IPMI output")
}

func SetFanSpeed(speed int) error {
	if speed < 0 || speed > 100 {
		return fmt.Errorf("invalid fan speed: %d (must be 0-100)", speed)
	}

	// Convert percentage to the value expected by IPMI
	// This formula may vary depending on your server
	ipmiValue := int(float64(speed) * 2.55) // 0-100% to 0-255

	cmd := exec.Command("ipmitool", "raw", "0x30", "0x45", "0x01", "0x00", fmt.Sprintf("0x%02x", ipmiValue))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to set fan speed to %d%%: %w, output: %s", speed, err, string(output))
	}

	log.Printf("Fan speed set to %d%%", speed)
	return nil
}
