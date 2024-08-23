package config

const (
	Version    = "v0.1.0"
	configFile = "config.json"
)

var (
	// Write default content to the file
	jsonTemplate = map[string]interface{}{
		"styles": map[string]string{
			"subtext": "#6c7086",
			"accent":  "#b4befe",
		},
		"items": []map[string]string{
			{"title": "Description : keybinding", "tag": "tag"},
			{"title": "Minimize Window : meta+m", "tag": "Kwin"},
		},
	}
)
