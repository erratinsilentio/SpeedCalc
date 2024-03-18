package selections

type options struct {
	DistanceUnit string
	SpeedUnit    string
	SpeedValue   string
}

var userSelections = map[string]options{
	"from": {
		DistanceUnit: "",
		SpeedUnit:    "",
		SpeedValue:   "",
	},
	"to": {
		DistanceUnit: "",
		SpeedUnit:    "",
		SpeedValue:   "",
	},
}

func SelectDistanceUnit(from, to string) {
	// Select to and from options
	fromOptions := userSelections["from"]
	toOptions := userSelections["to"]
	// Update the fields
	fromOptions.DistanceUnit = from
	toOptions.DistanceUnit = to
	// Assign the updated "to" options struct back to the map
	userSelections["from"] = fromOptions
	userSelections["to"] = toOptions
}

func SelectSpeedUnit(from, to string) {
	// Select to and from options
	fromOptions := userSelections["from"]
	toOptions := userSelections["to"]
	// Update the fields
	fromOptions.SpeedUnit = from
	toOptions.SpeedUnit = to
	// Assign the updated "to" options struct back to the map
	userSelections["from"] = fromOptions
	userSelections["to"] = toOptions
}

func SelectSpeedValue(from, to string) {
	// Select to and from options
	fromOptions := userSelections["from"]
	toOptions := userSelections["to"]
	// Update the fields
	fromOptions.SpeedValue = from
	toOptions.SpeedValue = to
	// Assign the updated "to" options struct back to the map
	userSelections["from"] = fromOptions
	userSelections["to"] = toOptions
}

func GetUserSelections() map[string]options {
	return userSelections
}
