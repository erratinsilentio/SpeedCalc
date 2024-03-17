package kmToKm

import (
	"fmt"

	"github.com/erratinsilentio/SpeedCalc/internal/selections"
	"github.com/manifoldco/promptui"
)

func GetPrompt() {

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
