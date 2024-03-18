package start

import (
	"fmt"

	"github.com/erratinsilentio/SpeedCalc/internal/calculations"
	"github.com/erratinsilentio/SpeedCalc/internal/prompts"
	"github.com/erratinsilentio/SpeedCalc/internal/selections"
	"github.com/manifoldco/promptui"
)

func SpeedCalc() {
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
		selections.SelectDistanceUnit("miles", "miles")
		prompts.MilesToMiles()
		prompts.SpeedValues()
		calculations.Convert()

	case "kilometers to kilometers":
		selections.SelectDistanceUnit("kilometers", "kilometers")
		prompts.KmToKm()
		prompts.SpeedValues()
		calculations.Convert()

	case "miles to kilometers":
		selections.SelectDistanceUnit("miles", "kilometers")
		prompts.MilesToKm()
		prompts.SpeedValues()
		calculations.Convert()

	case "kilometers to miles":
		selections.SelectDistanceUnit("kilometers", "miles")
		prompts.KmToMiles()
		prompts.SpeedValues()
		calculations.Convert()
		prompts.DisplayResult()
	}
}
