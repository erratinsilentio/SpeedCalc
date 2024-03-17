package kmToMiles

import (
	"fmt"

	"github.com/erratinsilentio/SpeedCalc/internal/selections"
	"github.com/manifoldco/promptui"
)

func GetPrompt() {

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
