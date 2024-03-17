package milesToMiles

import (
	"fmt"

	"github.com/erratinsilentio/SpeedCalc/internal/selections"
	"github.com/manifoldco/promptui"
)

func GetPrompt() {

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
