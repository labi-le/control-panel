package structures

// PanelSettings
// fields:
// Port - the port on which the panel will work
// Language - panel language
// Theme - panel theme
type PanelSettings struct {
	Port     string `json:"port"`
	Language string `json:"language"`
	Theme    string `json:"theme"`
}
