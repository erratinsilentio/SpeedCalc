package selections

type options struct {
	distanceUnit string
	speedUnit    string
	speedValue   string
}

var userSelections = map[string]options{
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

func SelectDistanceUnit(from, to string) {
	// Select to and from options
	fromOptions := userSelections["from"]
	toOptions := userSelections["to"]
	// Update the fields
	fromOptions.distanceUnit = from
	toOptions.distanceUnit = to
	// Assign the updated "to" options struct back to the map
	userSelections["from"] = fromOptions
	userSelections["to"] = toOptions
}

func SelectSpeedUnit(from, to string) {
	// Select to and from options
	fromOptions := userSelections["from"]
	toOptions := userSelections["to"]
	// Update the fields
	fromOptions.speedUnit = from
	toOptions.speedUnit = to
	// Assign the updated "to" options struct back to the map
	userSelections["from"] = fromOptions
	userSelections["to"] = toOptions
}

func SelectSpeedValue(from, to string) {
	// Select to and from options
	fromOptions := userSelections["from"]
	toOptions := userSelections["to"]
	// Update the fields
	fromOptions.speedValue = from
	toOptions.speedValue = to
	// Assign the updated "to" options struct back to the map
	userSelections["from"] = fromOptions
	userSelections["to"] = toOptions
}

func GetUserSelections() map[string]options {
	return userSelections
}
