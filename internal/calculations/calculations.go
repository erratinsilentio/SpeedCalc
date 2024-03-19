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
	if fromSpeedUnit == "miles/h" && toSpeedUnit == "min/mile" {
		result, err := milesPerHourToMinPerMile(fromSpeedValue)
		if err != nil {
			fmt.Println("Error while converting: %v", err)
			return
		}
		resultToString, err := floatToTimeString(result)
		if err != nil {
			fmt.Println("Error while converting float to time string")
		}
		selections.SetResult(resultToString)
		return
	}

	// min/mile to miles/h
	if fromSpeedUnit == "min/mile" && toSpeedUnit == "miles/h" {
		result, err := minPerMileToMilesPerHour(fromSpeedValue)
		if err != nil {
			fmt.Println("Error while converting: %v", err)
			return
		}
		resultToString := fmt.Sprintf("%.2f", result)
		selections.SetResult(resultToString)
		return
	}

	// km/h to min/km
	if fromSpeedUnit == "km/h" && toSpeedUnit == "min/km" {
		result, err := kmPerHourToMinPerKm(fromSpeedValue)
		if err != nil {
			fmt.Println("Error while converting: %v", err)
			return
		}
		resultToTime, _ := floatToTimeString(result)

		selections.SetResult(resultToTime)
		return
	}

	// min/km to km/h
	if fromSpeedUnit == "min/km" && toSpeedUnit == "km/h" {
		result, err := minPerKmToKmPerHour(fromSpeedValue)
		if err != nil {
			fmt.Println("Error while converting: %v", err)
			return
		}

		resultToString := fmt.Sprintf("%.2f", result)
		selections.SetResult(resultToString)
		return
	}

	// miles/h to km/h
	if fromSpeedUnit == "miles/h" && toSpeedUnit == "km/h" {
		result, err := milesPerHourToKmPerHour(fromSpeedValue)
		if err != nil {
			fmt.Println("Error while converting: %v", err)
			return
		}

		resultToString := fmt.Sprintf("%.2f", result)
		selections.SetResult(resultToString)
		return
	}

	// miles/h to min/km

	// min/miles to min/km

	// min/miles to km/h

	// km/h to miles/h
	if fromSpeedUnit == "km/h" && toSpeedUnit == "miles/h" {
		result, err := kmPerHourToMilePerHour(fromSpeedValue)
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

// CONVERSION FUNCTIONS

func milesPerHourToMinPerMile(s string) (f float64, e error) {
	milesPerHour, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.00, err
	}

	minutesPerMile := 60 / milesPerHour
	return minutesPerMile, nil
}

func minPerMileToMilesPerHour(s string) (f float64, e error) {
	totalSeconds, err := minutesToSeconds(s)
	if err != nil {
		return 0.0, err
	}

	minPerMile := float64(totalSeconds / 60)

	milesPerHour := 60 / minPerMile

	return milesPerHour, nil
}

func kmPerHourToMilePerHour(s string) (f float64, e error) {
	kmPerHour, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.00, err
	}

	conversionFactor := 0.621371192
	milesPerHour := kmPerHour * conversionFactor

	return milesPerHour, nil
}

func kmPerHourToMinPerKm(s string) (f float64, e error) {
	kmPerHour, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.00, err
	}

	minPerKm := 60 / kmPerHour
	return minPerKm, nil
}

func minPerKmToKmPerHour(s string) (f float64, e error) {
	totalSeconds, err := minutesToSeconds(s)
	if err != nil {
		return 0.0, err
	}

	minPerKm := float64(totalSeconds / 60)

	kmPerHour := 60 / minPerKm

	return kmPerHour, nil
}

func milesPerHourToKmPerHour(s string) (f float64, e error) {
	conversionRate := 1.60934
	milesPerHour, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Error: %v", err)
		return 0.00, err
	}

	kmPerHour := milesPerHour * conversionRate
	return kmPerHour, nil
}

// HELPER FUNCTIONS

func minutesToSeconds(timeString string) (float64, error) {
	if len(timeString) == 1 {
		minInt, _ := strconv.Atoi(timeString)
		return float64(minInt * 60), nil
	}

	parts := strings.Split(timeString, ":")

	if len(parts) != 2 {
		return 0, fmt.Errorf("Invalid time format")
	}

	minutes, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("Invalid minute format")
	}

	seconds, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("Invalid second format")
	}

	totalSeconds := float64(minutes*60 + seconds)

	return totalSeconds, nil
}

func floatToTimeString(f float64) (s string, e error) {
	minutes := int(f)
	seconds := (f - float64(minutes)) * 60

	minutesStr := strconv.Itoa(minutes)
	secondsStr := strconv.Itoa(int(seconds))

	if seconds < 10 {
		secondsStr = "0" + secondsStr
	}

	timeString := minutesStr + ":" + secondsStr

	return timeString, nil
}
