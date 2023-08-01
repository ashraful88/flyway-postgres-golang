package config

import "os"

type Config struct {
	APIPort         string
	FlywayLocations string
	FlywayConnRetry string
	FlywayDBSchemas string
}

var Conf Config

// return the config based on the environment
func GetConfigByEnv(env string) Config {
	var conf Config
	conf.APIPort = getEnvVarByName(env, "API_PORT")
	conf.FlywayConnRetry = getEnvVarByName(env, "FLYWAY_CONNECT_RETRIES")
	conf.FlywayLocations = getEnvVarByName(env, "FLYWAY_LOCATIONS")
	conf.FlywayDBSchemas = getEnvVarByName(env, "FLYWAY_SCHEMAS")

	return conf
}

// return the env variable based on the environment
func getEnvVarByName(env string, name string) string {
	envConfig := envConfigs[env][name]
	// reassign env variable if its in running environment variable or in .env file
	if os.Getenv(name) != "" {
		return os.Getenv(name)
	}

	return envConfig
}
