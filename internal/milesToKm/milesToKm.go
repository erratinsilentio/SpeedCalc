package milesToKm

import (
	"fmt"

	"github.com/erratinsilentio/SpeedCalc/internal/selections"
	"github.com/manifoldco/promptui"
)

func GetPrompt() {

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
