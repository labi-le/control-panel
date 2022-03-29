package structures

// PanelSettings
// Port - the port on which the panel will work
// Language - panel language
// Theme - panel theme
type PanelSettings struct {
	Port     string `json:"port"`
	Language string `json:"language"`
	Theme    string `json:"theme"`
}

func DefaultPanelSettings() *PanelSettings {
	return &PanelSettings{
		Port:     "7000",
		Language: "en",
		Theme:    "default",
	}
}
