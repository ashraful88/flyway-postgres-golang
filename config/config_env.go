package config

type kv map[string]string

// write configs here, not for keeping credentials
var envConfigs = map[string]kv{
	"local": {
		"API_PORT":               "8088",
		"FLYWAY_CONNECT_RETRIES": "5",
		"FLYWAY_LOCATIONS":       "filesystem:/flyway/sql",
		"FLYWAY_SCHEMAS":         "apps",
	},
	"staging": {
		"API_PORT":               "80",
		"FLYWAY_CONNECT_RETRIES": "60",
		"FLYWAY_LOCATIONS":       "filesystem:/app/sql",
		"FLYWAY_SCHEMAS":         "apps",
	},
	"production": {
		"API_PORT":               "80",
		"FLYWAY_CONNECT_RETRIES": "60",
		"FLYWAY_LOCATIONS":       "filesystem:/app/sql",
		"FLYWAY_SCHEMAS":         "apps",
	},
}
