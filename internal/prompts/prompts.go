package prompts

import (
	"fmt"

	"github.com/erratinsilentio/SpeedCalc/internal/calculations"
	"github.com/erratinsilentio/SpeedCalc/internal/selections"
	"github.com/manifoldco/promptui"
)

func MilesToMiles() {

	prompt := promptui.Select{
		Label: "And...",
		Items: []string{"from miles/h to min/mile", "from min/mile to miles/h"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Oh no! Looks like there is an error: %v", err)
		return
	}

	switch result {
	case "from miles/h to min/mile":
		selections.SelectSpeedUnit("miles/h", "min/mile")
		//todo
	case "from min/mile to miles/h":
		selections.SelectSpeedUnit("min/mile", "miles/h")
		//todo
	}
}

func KmToKm() {

	prompt := promptui.Select{
		Label: "And...",
		Items: []string{"from km/h to min/km", "from min/km to km/h"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Oh no! Looks like there is an error: %v", err)
		return
	}

	switch result {
	case "from km/h to min/km":
		selections.SelectSpeedUnit("km/h", "min/km")
		//todo
	case "from min/km to km/h":
		selections.SelectSpeedUnit("min/km", "km/h")
		//todo
	}
}

func MilesToKm() {

	prompt := promptui.Select{
		Label: "And...",
		Items: []string{
			"from miles/h to km/h",
			"from miles/h min/km",
			"from min/miles to min/km",
			"from min/miles to km/h",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Oh no! Looks like there is an error: %v", err)
		return
	}

	switch result {
	case "from miles/h to km/h":
		selections.SelectSpeedUnit("miles/h", "km/h")
		//todo
	case "from km/h to miles/h":
		selections.SelectSpeedUnit("km/h", "miles/h")
		//todo
	case "from min/miles to min/km":
		selections.SelectSpeedUnit("min/miles", "min/km")
		//todo
	case "from min/miles to km/h":
		selections.SelectSpeedUnit("min/miles", "km/h")
		//todo
	}
}

func KmToMiles() {

	prompt := promptui.Select{
		Label: "And...",
		Items: []string{
			"from km/h to miles/h",
			"from km/h to min/miles",
			"from min/km to miles/h",
			"from min/km to min/miles",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Oh no! Looks like there is an error: %v", err)
		return
	}

	switch result {
	case "from km/h to miles/h":
		selections.SelectSpeedUnit("km/h", "miles/h")
		//todo
	case "from km/h to min/miles":
		selections.SelectSpeedUnit("km/h", "min/miles")
		//todo
	case "from min/km to miles/h":
		selections.SelectSpeedUnit("min/km", "miles/h")
		//todo
	case "from min/km to min/miles":
		selections.SelectSpeedUnit("min/km", "min/miles")
		//todo
	}
}

func SpeedValues() {
	userSelections := selections.GetUserSelections()
	fromSpeedUnit := userSelections["from"].SpeedUnit

	prompt := promptui.Prompt{
		Label: fmt.Sprintf("enter value in %s", fromSpeedUnit),
	}

	fromValueString, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed, %v", err)
	}

	isPromptValid, message := calculations.CheckIfPromptIsValid("from", fromValueString)
	if !isPromptValid {
		fmt.Println(message)
	}

	selections.SelectSpeedValue(fromValueString)
}

func DisplayResult() {
	userSelections := selections.GetUserSelections()
	result := userSelections["to"].SpeedValue

	fmt.Printf("Result: %s\n", result)
}
