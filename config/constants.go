package config

const (
	Version    = "v0.1.1"
	configFile = "config.json"
)

var (
	// Default New user config.json
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
