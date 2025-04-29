package pkg

import (
	"log"
	"time"
)

const (
	WarningTemp  = 75.0 // Temperature threshold for berserk mode (100% fan)
	CriticalTemp = 80.0 // Temperature threshold for auto mode
)

var (
	isMonitoring    = false
	manualOverride  = false // Flag to track if user has manually overridden auto switch
	autoSwitchCount = 0     // Counter for auto switches (for logging)
)

// StartTemperatureMonitor begins monitoring the system temperature
func StartTemperatureMonitor() {
	if isMonitoring {
		return
	}

	isMonitoring = true
	go monitorLoop()
}

// SetManualOverride allows overriding the automatic mode switching
func SetManualOverride(override bool) {
	manualOverride = override
}

// monitorLoop continuously monitors temperature and adjusts fan settings
func monitorLoop() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	log.Println("Temperature monitoring service started")

	for isMonitoring {
		select {
		case <-ticker.C:
			temp, err := GetCPUTemperatures()
			if err != nil {
				log.Printf("Error reading temperature: %v", err)
				continue
			}

			// Check temperature thresholds and take action if needed
			if temp >= CriticalTemp {
				// Switch to auto mode when temperature is critical
				currentMode, err := GetFanMode()
				if err != nil {
					log.Printf("Error getting fan mode: %v", err)
					continue
				}

				if currentMode != "auto" && !manualOverride {
					log.Printf("CRITICAL TEMPERATURE ALERT: %.1f°C! Switching to auto mode", temp)
					err := SetFanMode("auto")
					if err != nil {
						log.Printf("Failed to switch to auto mode: %v", err)
					} else {
						autoSwitchCount++
						log.Printf("Auto-switched to auto mode (%d times) due to critical temperature", autoSwitchCount)
					}
				}
			} else if temp >= WarningTemp {
				// Switch to berserk (100% fan) when temperature is high
				currentMode, err := GetFanMode()
				if err != nil {
					log.Printf("Error getting fan mode: %v", err)
					continue
				}

				if currentMode == "manual" {
					currentSpeed, err := GetFanSpeed()
					if err != nil {
						log.Printf("Error getting fan speed: %v", err)
						continue
					}

					if currentSpeed < 100 && !manualOverride {
						log.Printf("HIGH TEMPERATURE WARNING: %.1f°C! Setting fans to 100%%", temp)
						err := SetFanSpeed(100)
						if err != nil {
							log.Printf("Failed to set fan speed to 100%%: %v", err)
						}
					}
				}
			}
		}
	}
}
