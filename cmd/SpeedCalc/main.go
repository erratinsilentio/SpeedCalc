package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type options struct {
	distanceUnit string
	speedUnit    string
	speedValue   string
}

func main() {

	userSelections := map[string]options{
		"from": {
			distanceUnit: "",
			speedUnit:    "",
			speedValue:   "",
		},
		"to": {
			distanceUnit: "",
			speedUnit:    "",
			speedValue:   "",
		},
	}

	fmt.Println("⚡️Welcome to SpeedCalc! Calculate running speed to make your training easier.")

	prompt := promptui.Select{
		Label: "Do we convert with miles or kilometers?",
		Items: []string{"miles to miles", "kilometers to kilometers", "miles to kilometers", "kilometers to miles"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Oh no! Looks like there is an error: %v", err)
		return
	}

	switch result {
	case "miles to miles":
		// Select to and from options
		fromOptions := userSelections["from"]
		toOptions := userSelections["to"]
		// Update the fields
		fromOptions.distanceUnit = "miles"
		toOptions.distanceUnit = "miles"
		// Assign the updated "to" options struct back to the map
		userSelections["from"] = fromOptions
		userSelections["to"] = toOptions

	case "kilometers to kilometers":
		// Select to and from options
		fromOptions := userSelections["from"]
		toOptions := userSelections["to"]
		// Update the fields
		fromOptions.distanceUnit = "kilometers"
		toOptions.distanceUnit = "kilometers"
		// Assign the updated "to" options struct back to the map
		userSelections["from"] = fromOptions
		userSelections["to"] = toOptions

	case "miles to kilometers":
		// Select to and from options
		fromOptions := userSelections["from"]
		toOptions := userSelections["to"]
		// Update the fields
		fromOptions.distanceUnit = "miles"
		toOptions.distanceUnit = "kilometers"
		// Assign the updated "to" options struct back to the map
		userSelections["from"] = fromOptions
		userSelections["to"] = toOptions

	case "kilometers to miles":
		// Select to and from options
		fromOptions := userSelections["from"]
		toOptions := userSelections["to"]
		// Update the fields
		fromOptions.distanceUnit = "kilometers"
		toOptions.distanceUnit = "miles"
		// Assign the updated "to" options struct back to the map
		userSelections["from"] = fromOptions
		userSelections["to"] = toOptions
	}
}
