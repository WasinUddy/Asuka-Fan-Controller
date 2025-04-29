package pkg

import (
	"fmt"
	"math"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func GetCPUTemperatures() (float64, error) {
	cmd := exec.Command("ipmitool", "sdr", "type", "temperature")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("failed to get temperature data: %w, output: %s", err, string(output))
	}

	// Parse the output to find temperature entries
	lines := strings.Split(string(output), "\n")
	var tempReadings []float64

	for _, line := range lines {
		// Look for Inlet Temp and Temp entries
		if (strings.Contains(line, "Inlet Temp") || strings.Contains(line, "Temp")) &&
			strings.Contains(line, "degrees C") {
			// Extract the temperature value using regex
			re := regexp.MustCompile(`\|\s+(\d+(?:\.\d+)?)\s+\|`)
			matches := re.FindStringSubmatch(line)
			if len(matches) >= 2 {
				temp, err := strconv.ParseFloat(matches[1], 64)
				if err == nil {
					tempReadings = append(tempReadings, temp)
				}
			}
		}
	}

	if len(tempReadings) == 0 {
		return 0, fmt.Errorf("no temperature data found")
	}

	// Calculate average temperature
	var sum float64
	for _, temp := range tempReadings {
		sum += temp
	}
	avgTemp := sum / float64(len(tempReadings))

	return math.Round(avgTemp*10) / 10, nil // Round to 1 decimal place
}
