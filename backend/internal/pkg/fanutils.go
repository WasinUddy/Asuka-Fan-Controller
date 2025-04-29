package pkg

import (
	"fmt"
	"github.com/WasinUddy/Ayanami-Fan-Controller/internal/utils"
	"math"
	"strconv"
)

func SetFanMode(mode string) error {
	if mode != "auto" && mode != "manual" {
		return fmt.Errorf("invalid mode: %s, must be 'auto' or 'manual'", mode)
	}

	var cmdArgs []string
	if mode == "auto" {
		cmdArgs = []string{"raw", "0x30", "0x30", "0x01", "0x01"}
	} else {
		cmdArgs = []string{"raw", "0x30", "0x30", "0x01", "0x00"}
	}

	return utils.RunIpmiCommand(cmdArgs...)
}

func SetFanSpeed(speed int) error {
	if speed < 0 || speed > 100 {
		return fmt.Errorf("invalid speed: %d, must be between 0 and 100", speed)
	}

	/*
		Percentage	Hex Value	RPM (Approximate, varies by fan)
		0%	0x00	Fans off (not recommended)
		5%	0x05	Very low speed
		10%	0x0a	Low speed
		15%	0x0f	Quiet but may be insufficient
		20%	0x14	Common for low noise
		25%	0x19	Balanced for noise/cooling
		30%	0x1e	Moderate cooling
		40%	0x28	Good cooling
		50%	0x32	Medium-high speed
		60%	0x3c	High speed
		75%	0x4b	Very high speed
		100%	0x64	Maximum speed

		Thank you grok AI
	*/
	roundedSpeed := int(math.Round(float64(speed)/5.0)) * 5

	hexSpeed := fmt.Sprintf("0x%02x", roundedSpeed)

	parsedHexSpeed, err := strconv.ParseUint(hexSpeed, 0, 8)
	if err != nil {
		return fmt.Errorf("failed to parse hex speed: %v", err)
	}
	hexSpeedStr := fmt.Sprintf("0x%02x", parsedHexSpeed)

	cmdArgs := []string{"raw", "0x30", "0x30", "0x02", "0xff", hexSpeedStr}

	return utils.RunIpmiCommand(cmdArgs...)
}
