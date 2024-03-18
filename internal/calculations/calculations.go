package calculations

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/erratinsilentio/SpeedCalc/internal/selections"
)

func CheckIfPromptIsValid(u, s string) bool {
	userSelections := selections.GetUserSelections()

	speedUnit := userSelections[u].SpeedUnit

	switch speedUnit {
	case "miles/h":
		//todo
	case "min/mile":
		//todo
	case "km/h":
		//todo
	case "min/km":
		//todo
	}
	// Regular expression for matching time format like "6:30"
	timeFormatRegex := regexp.MustCompile(`^\d+:\d+$`)

	// Check if the string matches the time format
	if timeFormatRegex.MatchString(s) {
		parts := strings.Split(s, ":")
		if len(parts) == 2 {
			minutes, err := strconv.Atoi(parts[0])
			if err != nil {
				return false
			}
			seconds, err := strconv.Atoi(parts[1])
			if err != nil {
				return false
			}
			if minutes < 0 || seconds < 0 || seconds >= 60 {
				return false
			}
			return true
		}
	}

	// Try to convert the string into a float
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}
	return true
}

func StringToFloat(s string) float64 {
	//todo
	return 0.00
}

func Convert() {
	// check from and to units
	// convert strings to floats
	// perform conversion
}
