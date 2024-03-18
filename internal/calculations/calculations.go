package calculations

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/erratinsilentio/SpeedCalc/internal/selections"
)

func checkIfTimeValid(s string) (x bool, y string) {
	// Regular expression for matching time format like "6:30"
	timeFormatRegex := regexp.MustCompile(`^\d+:\d+$`)
	// Error message to display when input invalid
	errorMessage := "Invalid input, follow '6:30' format for time values!"
	// Check if the string matches the time format
	if timeFormatRegex.MatchString(s) {
		parts := strings.Split(s, ":")
		if len(parts) == 2 {
			minutes, err := strconv.Atoi(parts[0])
			if err != nil {
				return false, errorMessage
			}
			seconds, err := strconv.Atoi(parts[1])
			if err != nil {
				return false, errorMessage
			}
			if minutes < 0 || seconds < 0 || seconds >= 60 {
				return false, errorMessage
			}
			return true, "Success"
		}
	}
	return false, errorMessage
}

func checkIfNumberValid(s string) (x bool, y string) {
	// Try to convert the string into a float
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false, "Value invalid, use plain numbers like '9'"
	}
	return true, "Success"
}

func CheckIfPromptIsValid(u, s string) (x bool, y string) {
	userSelections := selections.GetUserSelections()

	speedUnit := userSelections[u].SpeedUnit

	switch speedUnit {
	case "miles/h":
		valid, message := checkIfNumberValid(s)
		return valid, message
	case "min/mile":
		valid, message := checkIfTimeValid(s)
		return valid, message
	case "km/h":
		valid, message := checkIfNumberValid(s)
		return valid, message
	case "min/km":
		valid, message := checkIfTimeValid(s)
		return valid, message
	}

	return false, "Unknown error!"
}

func StringToFloat(speedUnit, speedValue string) float64 {
	switch speedValue {
	case "km/h":
		//todo
	case "min/km":
		//todo
	case "miles/h":
		//todo
	case "min/mile":
		//todo
	}

	return 0.00
}

func Convert() {
	// check from and to units
	userSelections := selections.GetUserSelections()

	fromSpeedUnit := userSelections["from"].SpeedUnit
	fromSpeedValue := userSelections["from"].SpeedValue

	toSpeedUnit := userSelections["to"].SpeedUnit

	// miles/h to min/mile

	// min/mile to miles/h

	// km/h to min/km

	// min/km to km/h

	// miles/h to km/h

	// miles/h to min/km

	// min/miles to min/km

	// min/miles to km/h

	// km/h to miles/h
	if fromSpeedUnit == "km/h" && toSpeedUnit == "miles/h" {
		result, err := KmPerHourToMilePerHour(fromSpeedValue)
		if err != nil {
			fmt.Println("Error while converting: %v", err)
			return
		}
		resultToString := fmt.Sprintf("%.2f", result)
		selections.SetResult(resultToString)
		return
	}

	// km/h to min/miles

	// min/km to miles/h

	// min/km to min/miles

}

func KmPerHourToMilePerHour(s string) (f float64, e error) {
	kmPerHour, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.00, err
	}

	conversionFactor := 0.621371192
	milesPerHour := kmPerHour * conversionFactor

	return milesPerHour, nil
}
