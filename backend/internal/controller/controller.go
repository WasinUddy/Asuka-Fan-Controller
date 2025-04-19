package controller

import (
	"fmt"
	"github.com/WasinUddy/Asuka-Fan-Controller/internal/pkg"
	"net/http"
	"strconv"
	"sync"
)

var (
	fanMode  = "auto"
	fanSpeed = 50
	fanLock  = sync.RWMutex{}
)

func GetFanStatus(w http.ResponseWriter, r *http.Request) {
	fanLock.RLock()
	defer fanLock.RUnlock()

	// Return in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, `{"mode": "%s", "speed": %d}`, fanMode, fanSpeed)
	if err != nil {
		return
	}
}

func SetFanMode(w http.ResponseWriter, r *http.Request) {
	fanLock.Lock()
	defer fanLock.Unlock()

	mode := r.URL.Query().Get("mode")
	if mode != "auto" && mode != "manual" {
		http.Error(w, "Invalid mode must use 'auto' or 'manual'", http.StatusBadRequest)
		return
	}

	fanMode = mode

	var err error

	err = pkg.SetFanMode(fanMode)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to set fan mode: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "Fan mode set to %s", fanMode)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}

func SetFanSpeed(w http.ResponseWriter, r *http.Request) {
	fanLock.Lock()
	defer fanLock.Unlock()

	speed := r.URL.Query().Get("speed")
	if speed == "" {
		http.Error(w, "Speed parameter is required", http.StatusBadRequest)
		return
	}

	var err error
	fanSpeed, err = strconv.Atoi(speed)
	if err != nil || fanSpeed < 0 || fanSpeed > 100 {
		http.Error(w, "Invalid speed value, must be between 0 and 100", http.StatusBadRequest)
		return
	}

	err = pkg.SetFanSpeed(fanSpeed)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to set fan speed: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "Fan speed set to %d%%", fanSpeed)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
